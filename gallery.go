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
	Function       interface{}
	Args           []interface{}
	PanelDimension string
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

func (s *Server) RegisterComponent(name string, panelDimension string, fn interface{}, args ...interface{}) {
	s.components[name] = ComponentInfo{
		Name:           name,
		Function:       fn,
		Args:           args,
		PanelDimension: panelDimension,
	}

	s.componentOrder = append(s.componentOrder, name)
}

func (s *Server) generateHTML() string {
	page := ui.NewPage().
		SetTitle("konstfish/ui Gallery").
		AddStyleSheet("static/main.css").
		AddStyleSheet("static/gallery/etc.css").
		AddLinkWithType("image/svg+xml", "static/logo.svg", "icon").
		AddScript("https://unpkg.com/htmx.org@2.0.4")

	page.Body.AddChild(ui.NewElement("h1").SetContent("konstfish/ui Gallery"))

	componentGroup := ui.NewElement("div").AddClass("gallery-component-group")
	for _, name := range s.componentOrder {
		componentGroup.AddChild(
			ui.NewElement("fieldset").
				AddClass("panel").
				AddClass(s.components[name].PanelDimension).
				SetAttribute("id", fmt.Sprintf("comp-%s", name)).
				AddChild(ui.NewElement("span").
					SetAttribute("hx-get", fmt.Sprintf("/%s", name)).
					SetAttribute("hx-swap", "outerHTML").
					SetAttribute("hx-trigger", "load")),
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
	server.RegisterComponent("Text", "d1x1", kf.Text, "Hello, world!")
	server.RegisterComponent("Link", "d1x1", kf.Link, "Links", "https://github.com/konstfish/ui")
	server.RegisterComponent("Panel", "d1x1", kf.Panel, kf.Text("Panels"))

	server.RegisterComponent("Button", "d1x1", kf.Group, kf.Button("Button"), kf.Button("Button 2"))

	server.RegisterComponent("Spinner", "d1x1", kf.Spinner, "Loading...")
	server.RegisterComponent("Fieldset", "d1x2", kf.Fieldset, "Fieldset", kf.Text("Span in fieldset!"))

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
