package ui

import (
	"bytes"
	"text/template"
)

type Page struct {
	Lang string
	Head *Element
	Body *Element
}

func NewPage() *Page {
	head := NewElement("head")

	charsetMeta := NewElement("meta").SetAttribute("charset", "UTF-8")
	head.AddChild(charsetMeta)

	viewportMeta := NewElement("meta").
		SetAttribute("name", "viewport").
		SetAttribute("content", "width=device-width, initial-scale=1.0")
	head.AddChild(viewportMeta)

	return &Page{
		Lang: "en",
		Head: head,
		Body: NewElement("body"),
	}
}

func (p *Page) AddMeta(name, content string) *Page {
	meta := NewElement("meta").
		SetAttribute("name", name).
		SetAttribute("content", content)
	p.Head.AddChild(meta)
	return p
}

func (p *Page) SetTitle(title string) *Page {
	titleElement := NewElement("title").SetContent(title)
	p.Head.AddChild(titleElement)
	return p
}

func (p *Page) SetDescription(desc string) *Page {
	return p.AddMeta("description", desc)
}

func (p *Page) AddLink(rel string, href string) *Page {
	link := NewElement("link").
		SetAttribute("rel", rel).
		SetAttribute("href", href)
	p.Head.AddChild(link)
	return p
}

func (p *Page) AddLinkWithType(rel string, href string, linkType string) *Page {
	link := NewElement("link").
		SetAttribute("rel", rel).
		SetAttribute("href", href).
		SetAttribute("type", linkType)
	p.Head.AddChild(link)
	return p
}

func (p *Page) AddStyleSheet(href string) *Page {
	link := NewElement("link").
		SetAttribute("rel", "stylesheet").
		SetAttribute("href", href)
	p.Head.AddChild(link)
	return p
}

func (p *Page) AddScript(src string) *Page {
	script := NewElement("script").SetAttribute("src", src)
	p.Head.AddChild(script)
	return p
}

var pageTemplate = template.Must(TemplateElement.New("page").Parse(`
<!DOCTYPE html>
<html lang="{{.Lang}}">
{{template "element" .Head}}
{{template "element" .Body}}
</html>`))

func (p *Page) Render() (string, error) {
	var buf bytes.Buffer
	err := pageTemplate.Execute(&buf, p)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
