package main

import (
	"github.com/fipress/GoUI"
	"github.com/fipress/GoUI/widgets/filepicker"
)

func main() {
	//goui.Service("chat/:msg", chatService)
	goui.RegisterWidgets(new(filepicker.FilePicker))

	goui.Create(goui.Settings{Title: "FilePicker", Top: 30, Left: 100, Width: 300, Height: 440})
}
