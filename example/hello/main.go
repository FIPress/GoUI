package main

import (
	"github.com/fipress/GoUI"
)

var on = true

func getMsg() string {
	if on {
		on = false
		return "Hello world :-)"
	} else {
		on = true
		return "And hello dear developer :-)"
	}
}

func main() {
	//register a service
	goui.Service("hello", func(context *goui.Context) {
		context.Success(getMsg())
	})

	//create and open a window
	goui.Create(goui.Settings{Title: "Hello",
		Left:      200,
		Top:       50,
		Width:     400,
		Height:    510,
		Resizable: true,
		Debug:     true})
}
