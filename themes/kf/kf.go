package kf

import ui "github.com/konstfish/ui/core"

func Panel(text string) (string, error) {
	panel := ui.NewElement("div").
		AddClass("panel").
		AddClass("panel-adjust")

	panel.SetContent(text)

	return panel.Render()
}

func Spinner(text string) (string, error) {
	spinner := ui.NewElement("span")

	spinner.AddChild(ui.NewElement("span").AddClass("spinner"))

	if text != "" {
		spinner.SetContent(" " + text)
	}

	return spinner.Render()
}
