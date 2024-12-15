package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/konstfish/ui/themes/kf"
)

type ComponentInfo struct {
	Name     string
	Function interface{}
	Args     []interface{}
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

func (s *Server) RegisterComponent(group string, name string, fn interface{}, args ...interface{}) {
	if s.components[group] == nil {
		s.components[group] = make(map[string]ComponentInfo)
	}
	s.components[group][name] = ComponentInfo{
		Name:     name,
		Function: fn,
		Args:     args,
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
	<style>
	fieldset {
		margin: 1rem;
		padding: 1rem;
		border: 1px solid var(--bg-sec);
		display: grid;
		align-items: center;
		justify-content: center;
		width: 100px;
		height: 100px;
		border-radius: var(--border-rad-main);
	}
    legend{
        margin-bottom: -8%; /* magic number type shit */
    }
	</style>
</head>
<body>
    <script src="https://unpkg.com/htmx.org@2.0.4"></script>

	<h1>konstfish/ui Gallery</h1>
`

	for group := range s.components {
		html += fmt.Sprintf("  <h2>%s</h2>\n", group)
		html += "  <div class=\"component-group\">"
		for name := range s.components[group] {
			html += fmt.Sprintf("  <fieldset><legend>%s</legend>\n", name)
			html += fmt.Sprintf("    <span hx-get=\"/%s/%s\" hx-swap=\"outerHTML\" hx-trigger=\"load\"></span>\n", group, name)
			html += "  </fieldset>\n"
		}
		html += "  </div>\n"
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

		if len(results) == 2 && !results[1].IsNil() {
			http.Error(w, results[1].Interface().(error).Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, results[0].String())
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
	server.RegisterComponent("Base", "Panel", kf.Panel, "Hello, world!")

	server.RegisterComponent("Extras", "Spinner", kf.Spinner, "Loading...")

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
