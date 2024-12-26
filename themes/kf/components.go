package kf

import (
	"fmt"

	ui "github.com/konstfish/ui/core"
)

// Group creates a div element with the given content.
func Group(content ...*ui.Element) *ui.Element {
	panel := ui.NewElement("div")

	for _, c := range content {
		panel.AddChild(c)
	}

	return panel
}

// GroupClass creates a div element with the given class and content.
func GroupClass(class string, content ...*ui.Element) *ui.Element {
	panel := ui.NewElement("div").AddClass(class)

	for _, c := range content {
		panel.AddChild(c)
	}

	return panel
}

// Panel creates a div element with the given content & the "panel" class.
func Panel(content *ui.Element) *ui.Element {
	panel := ui.NewElement("div").
		AddClass("panel").
		AddChild(content)

	return panel
}

// Text creates a span element with the given text content.
func Text(text string) *ui.Element {
	return ui.NewElement("span").SetContent(text)
}

// Link creates an anchor element with the given text content and URL.
func Link(text string, url string) *ui.Element {
	link := ui.NewElement("a").
		SetAttribute("href", url).
		SetContent(text)

	return link
}

// Fieldset creates a fieldset element with the given legend and content.
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
	return ui.NewElement("div").
		AddClass("header-title").
		SetContent(text).
		AddChild(ui.NewElement("img").SetAttribute("src", logoSrc).AddClass("icon"))
}

// Code creates a code block with syntax highlighting. This depends on the Prism.js library.
func Code(lang string, code string) *ui.Element {
	return ui.NewElement("pre").AddChild(ui.NewElement("code").SetAttribute("class", fmt.Sprintf("language-%s", lang)).SetContent(code))
}
