package main

import (
	"github.com/fipress/GoUI"
)

func main() {
	goui.Service("hello", func(context *goui.Context) {
		context.Success("Hello world!")
	})

	goui.Create(goui.Settings{"Hello", "./ui/hello.html", 20, 30, 300, 200, true, true})
}
