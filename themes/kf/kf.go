package kf

import ui "github.com/konstfish/ui/core"

func Group(content ...*ui.Element) *ui.Element {
	panel := ui.NewElement("div")

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

func Button(text string) *ui.Element {
	return ui.NewElement("button").SetContent(text)
}

func Fieldset(legend string, content *ui.Element) *ui.Element {
	fieldset := ui.NewElement("fieldset").
		AddClass("panel").
		AddClass("panel-adjust").
		AddChild(ui.NewElement("legend").SetContent(legend)).
		AddChild(content)

	return fieldset
}

func Spinner(text string) *ui.Element {
	spinner := ui.NewElement("span")

	spinner.AddChild(ui.NewElement("span").AddClass("spinner"))

	if text != "" {
		spinner.SetContent(" " + text)
	}

	return spinner
}
