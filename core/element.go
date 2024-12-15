package ui

import (
	"bytes"
	"html"
	"text/template"
)

type Element struct {
	Tag        string
	Classes    []string
	Attributes map[string]string
	Content    interface{}
	Children   []*Element
}

func NewElement(tag string) *Element {
	return &Element{
		Tag:        tag,
		Classes:    make([]string, 0),
		Attributes: make(map[string]string),
		Children:   make([]*Element, 0),
	}
}

func (e *Element) AddClass(class string) *Element {
	e.Classes = append(e.Classes, class)
	return e
}

func (e *Element) SetAttribute(key, value string) *Element {
	e.Attributes[key] = value
	return e
}

func (e *Element) SetContent(content string) *Element {
	e.Content = html.EscapeString(content)
	return e
}

func (e *Element) AddChild(child *Element) *Element {
	e.Children = append(e.Children, child)
	return e
}

var tmpl = template.Must(template.New("element").Parse(`
{{define "element"}}<{{.Tag}}{{if .Classes}} class="{{range $i, $class := .Classes}}{{if $i}} {{end}}{{$class}}{{end}}"{{end}}{{range $key, $value := .Attributes}} {{$key}}="{{$value}}"{{end}}>{{range .Children}}{{template "element" .}}{{end}}{{if .Content}}{{.Content}}{{end}}</{{.Tag}}>{{end}}
`))

func (e *Element) Render() (string, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "element", e)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
