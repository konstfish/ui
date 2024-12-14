package main

import (
	"fmt"

	ui "github.com/konstfish/ui/core"
)

func main() {
	div := ui.NewElement("div").
		AddClass("container").
		SetAttribute("id", "main").
		AddChild(
			ui.NewElement("h1").
				AddClass("title").
				SetContent("test"),
		).
		AddChild(
			ui.NewElement("p").
				AddClass("content").
				SetContent("asdf"),
		)
	out, err := div.Render()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(out)

	xss := ui.NewElement("div").
		SetContent("<script>alert('xss')</script>")

	out2, err := xss.Render()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println(out2)
}
