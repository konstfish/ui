package kf

import (
	"fmt"

	ui "github.com/konstfish/ui/core"
)

func Group(content ...*ui.Element) *ui.Element {
	panel := ui.NewElement("div")

	for _, c := range content {
		panel.AddChild(c)
	}

	return panel
}

func GroupClass(class string, content ...*ui.Element) *ui.Element {
	panel := ui.NewElement("div").AddClass(class)

	for _, c := range content {
		panel.AddChild(c)
	}

	return panel
}

func Panel(content *ui.Element) *ui.Element {
	panel := ui.NewElement("div").
		AddClass("panel").
		AddClass("panel-adjust").
		AddChild(content)

	return panel
}

func Text(text string) *ui.Element {
	return ui.NewElement("span").SetContent(text)
}

func Link(text string, url string) *ui.Element {
	link := ui.NewElement("a").
		SetAttribute("href", url).
		SetContent(text)

	return link
}

func Fieldset(legend string, content *ui.Element) *ui.Element {
	fieldset := ui.NewElement("fieldset").
		AddClass("panel").
		AddClass("panel-adjust").
		AddChild(ui.NewElement("legend").SetContent(legend)).
		AddChild(content)

	return fieldset
}

func Title(text string) *ui.Element {
	return ui.NewElement("h1").SetContent(text)
}

func TitleLogo(text string, logoSrc string) *ui.Element {
	return ui.NewElement("header-text").
		SetContent(text).
		AddChild(ui.NewElement("img").SetAttribute("src", logoSrc).AddClass("icon"))
}

func Header(logo *ui.Element, links []KeyValue) *ui.Element {
	var headerElement *ui.Element = ui.NewElement("header").AddChild(logo)

	var linksElement *ui.Element
	if links != nil {
		linksElement = ui.NewElement("nav")
		for _, kv := range links {
			linksElement.AddChild(Link(kv.Key, kv.Value))
		}

		headerElement.AddChild(linksElement)
	}

	return headerElement
}

func SeparatorBar() *ui.Element {
	return ui.NewElement("div").AddClass("separator-bar")
}

/* depends on prism */
func Code(lang string, code string) *ui.Element {
	return ui.NewElement("pre").AddChild(ui.NewElement("code").SetAttribute("class", fmt.Sprintf("language-%s", lang)).SetContent(code))
}
