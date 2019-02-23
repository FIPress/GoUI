package main

import (
	"github.com/fipress/GoUI"
)

func main() {
	//register a service
	goui.Service("hello", func(context *goui.Context) {
		context.Success("Hello world!")
	})

	//create and open a window
	goui.Create(goui.Settings{Title: "Hello",
		Url:       "./ui/hello.html",
		Left:      20,
		Top:       30,
		Width:     300,
		Height:    200,
		Resizable: true,
		Debug:     true})
}
