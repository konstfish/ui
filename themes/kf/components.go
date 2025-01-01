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
func Panel(content ...*ui.Element) *ui.Element {
	panel := ui.NewElement("div").
		AddClass("panel")

	for _, c := range content {
		panel.AddChild(c)
	}

	return panel
}

// Text creates a span element with the given text content.
func Text(text string) *ui.Element {
	return ui.NewElement("span").SetContent(text)
}

// Paragraph creates a p element with the given text content.
func Paragraph(text string) *ui.Element {
	return ui.NewElement("p").SetContent(text)
}

// Header1 creates an h1 element with the given text content.
func Header1(text string) *ui.Element {
	return ui.NewElement("h1").SetContent(text)
}

// Header2 creates an h2 element with the given text content.
func Header2(text string) *ui.Element {
	return ui.NewElement("h2").SetContent(text)
}

// Header3 creates an h3 element with the given text content.
func Header3(text string) *ui.Element {
	return ui.NewElement("h3").SetContent(text)
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

func Form() *ui.Element {
	return ui.NewElement("form")
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

func List(items ...*ui.Element) *ui.Element {
	list := ui.NewElement("ul")

	for _, item := range items {
		list.AddChild(ui.NewElement("li").AddChild(item))
	}

	return list
}
