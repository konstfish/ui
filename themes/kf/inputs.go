package kf

import ui "github.com/konstfish/ui/core"

func Button(text string) *ui.Element {
	return ui.NewElement("button").SetContent(text)
}

func ButtonDanger(text string) *ui.Element {
	return Button(text).AddClass("danger")
}

func ButtonIcon(text string, iconSource string) *ui.Element {
	return Button(text).AddChild(ui.NewElement("img").SetAttribute("src", iconSource).AddClass("icon"))
}

func Input(placeholder string) *ui.Element {
	return ui.NewElement("input").
		SetAttribute("type", "text").
		SetAttribute("placeholder", placeholder)
}

func Dropdown(options []string) *ui.Element {
	dropdown := ui.NewElement("select")

	for _, option := range options {
		dropdown.AddChild(ui.NewElement("option").SetContent(option))
	}

	return dropdown
}
