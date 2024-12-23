package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"reflect"

	ui "github.com/konstfish/ui/core"
	"github.com/konstfish/ui/themes/kf"
)

type ComponentInfo struct {
	Name           string
	FunctionString string
	Function       interface{}
	Args           []interface{}
}

type Server struct {
	componentOrder []string
	components     map[string]ComponentInfo
	port           string
}

func NewServer(port string) *Server {
	return &Server{
		components: make(map[string]ComponentInfo),
		port:       port,
	}
}

func (s *Server) RegisterComponent(name string, functionString string, fn interface{}, args ...interface{}) {
	s.components[name] = ComponentInfo{
		Name:           name,
		Function:       fn,
		FunctionString: functionString,
		Args:           args,
	}

	s.componentOrder = append(s.componentOrder, name)
}

func (s *Server) generateHTML() string {
	page := ui.NewPage().
		SetTitle("konstfish/ui Gallery").
		AddScript("https://unpkg.com/htmx.org@2.0.4").
		AddScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/prism.min.js").
		AddScript("https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/components/prism-go.min.js").
		AddStyleSheet("static/main.css").
		AddStyleSheet("static/prism.css").
		AddStyleSheet("static/gallery/etc.css").
		AddLinkWithType("icon", "static/logo.svg", "image/svg+xml")

	// page.Body.AddChild(ui.NewElement("h1").SetContent("konstfish/ui Gallery").AddChild(ui.NewElement("img").SetAttribute("src", "static/logo.svg").AddClass("logo").AddClass("icon")))

	page.Body.AddChild(kf.Header(kf.TitleLogo("konstfish/ui Gallery", "static/logo.svg"), []kf.KeyValue{{"Source", "https://github.com/konstfish/ui"}, {"Docs", "https://pkg.go.dev/github.com/konstfish/ui/core"}}))

	/*page.Body.AddChild(
		kf.GroupClass("gallery-header", kf.ButtonIcon("Source", "https://upload.wikimedia.org/wikipedia/commons/9/91/Octicons-mark-github.svg").SetAttribute("onclick", "location.href='http://github.com/konstfish/ui'")),
	)*/

	componentGroup := ui.NewElement("div").AddClass("gallery-component-group")
	for _, name := range s.componentOrder {
		componentGroup.AddChild(
			ui.NewElement("div").
				AddClass("gallery-component").
				SetAttribute("id", fmt.Sprintf("comp-%s", name)).
				AddChild(kf.Group().AddClass("flip").AddClass("panel").
					AddChild(kf.Group(kf.Placeholder(fmt.Sprintf("/%s", name))).AddClass("flip-front")).
					AddChild(kf.Code("go", s.components[name].FunctionString).AddClass("flip-back")),
				),
		)
	}

	page.Body.AddChild(componentGroup)

	html, err := page.Render()
	if err != nil {
		log.Fatal(err)
		return "whoops"
	}

	return html
}

func (s *Server) handleComponent(info ComponentInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn := reflect.ValueOf(info.Function)

		args := make([]reflect.Value, len(info.Args))
		for i, arg := range info.Args {
			args[i] = reflect.ValueOf(arg)
		}

		results := fn.Call(args)

		out, err := results[0].Interface().(*ui.Element).Render()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, out)
	}
}

//go:embed static/*
var staticFiles embed.FS

func (s *Server) Start() error {
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(staticFS))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, s.generateHTML())
			return
		}
		http.NotFound(w, r)
	})

	for _, info := range s.components {
		http.HandleFunc(fmt.Sprintf("/%s", info.Name), s.handleComponent(info))
	}

	fmt.Printf("Server starting on port %s\n", s.port)
	return http.ListenAndServe(s.port, nil)
}

func main() {
	server := NewServer(":8080")
	server.RegisterComponent("Text", "kf.Text(\"Hello, world!\")", kf.Text, "Hello, world!")
	server.RegisterComponent("Link", "d1x1", kf.Link, "Links", "https://github.com/konstfish/ui")
	server.RegisterComponent("Panel", "d1x1", kf.Panel, kf.Text("Panels"))

	server.RegisterComponent("Button", "d1x1", kf.GroupClass, "button-display", kf.Button("Button"), kf.ButtonDanger("Button Danger"))
	server.RegisterComponent("Input", "d1x1", kf.Input, "Placeholder")
	server.RegisterComponent("Dropdown", "d1x1", kf.Dropdown, []string{"Option 1", "Option 2", "Option 3"})

	server.RegisterComponent("Spinner", "d1x1", kf.Spinner, "Loading...")
	server.RegisterComponent("Fieldset", "d1x2", kf.Fieldset, "Fieldset", kf.Text("Span in fieldset!"))

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
