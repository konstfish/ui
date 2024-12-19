package main

import (
	"fmt"
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
	components map[string]map[string]ComponentInfo
	port       string
}

func NewServer(port string) *Server {
	return &Server{
		components: make(map[string]map[string]ComponentInfo),
		port:       port,
	}
}

func (s *Server) RegisterComponent(group string, name string, panelDimension string, fn interface{}, args ...interface{}) {
	if s.components[group] == nil {
		s.components[group] = make(map[string]ComponentInfo)
	}

	s.components[group][name] = ComponentInfo{
		Name:           name,
		Function:       fn,
		Args:           args,
		PanelDimension: panelDimension,
	}
}

func (s *Server) generateHTML() string {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>konstfish/ui</title>
    <link rel="stylesheet" href="static/main.css">
	<link rel="icon" type="image/svg+xml" href="static/logo.svg">
	<!-- gallery specifics -->
	<style>
		.gallery-component-group {
			display: flex;
			flex-wrap: wrap;
		}
		.d1x1 {
			width: 120px;
			height: 120px;
		}
		.d1x2 {
			width: calc(240px + var(--margin)*4 + 2px);
			height: 120px;
		}
	</style>
</head>
<body>
    <script src="https://unpkg.com/htmx.org@2.0.4"></script>

	<h1>konstfish/ui Gallery</h1>
`

	for group := range s.components {
		head, _ := ui.NewElement("h2").SetContent(group).Render()
		html += head

		componentGroup := ui.NewElement("div").AddClass("gallery-component-group")
		for name := range s.components[group] {
			componentGroup.AddChild(
				ui.NewElement("fieldset").
					AddClass("panel").
					AddClass(s.components[group][name].PanelDimension).
					SetAttribute("id", fmt.Sprintf("%s-%s", group, name)).
					AddChild(ui.NewElement("legend").
						SetContent(name)).
					AddChild(ui.NewElement("span").
						SetAttribute("hx-get", fmt.Sprintf("/%s/%s", group, name)).
						SetAttribute("hx-swap", "outerHTML").
						SetAttribute("hx-trigger", "load")),
			)
		}

		componentGroupOut, _ := componentGroup.Render()
		html += componentGroupOut
	}

	html += "</body>\n</html>"
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

func (s *Server) Start() error {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, s.generateHTML())
			return
		}
		http.NotFound(w, r)
	})

	for group, components := range s.components {
		for _, info := range components {
			http.HandleFunc("/"+group+"/"+info.Name, s.handleComponent(info))
		}
	}

	fmt.Printf("Server starting on port %s\n", s.port)
	return http.ListenAndServe(s.port, nil)
}

func main() {
	server := NewServer(":8080")
	server.RegisterComponent("Base", "Text", "d1x1", kf.Text, "Hello, world!")
	server.RegisterComponent("Base", "Link", "d1x1", kf.Link, "konstfish/ui", "https://github.com/konstfish/ui")
	server.RegisterComponent("Base", "Panel", "d1x1", kf.Panel, kf.Text("Panel content!"))

	server.RegisterComponent("Inputs", "Button", "d1x1", kf.Button, "Click me!")

	server.RegisterComponent("Extras", "Spinner", "d1x1", kf.Spinner, "Loading...")
	server.RegisterComponent("Extras", "Fieldset", "d1x2", kf.Fieldset, "Fieldset", kf.Text("Span in fieldset!"))

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
