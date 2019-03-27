package main

import (
	"github.com/fipress/GoUI"
	"strings"
)

const (
	greetings = "Hello, I'm your pal B. How may I help you today?"
	boast     = "You can ask me anything. There is nothing I don't know."
	truth     = "Uh-oh~~~  you got me! I'm just a dummy robot."
	//pickup = "Knock knock..."
	askfor = "Say something please..."
)

func main() {
	goui.Service("chat/:msg", chatService)

	goui.Create(goui.Settings{Title: "Chatbot", Top: 30, Left: 100, Width: 300, Height: 440})
}

func chatService(ctx *goui.Context) {
	msg := ctx.GetParam("msg")
	receive(msg)
}

func receive(msg string) {
	msg = strings.ToLower(msg)
	switch {
	case msg == "":
		send(askfor)
	case msg == "ready":
		send(greetings)
	case strings.HasSuffix(msg, "?"):
		send(truth)
	default:
		send(boast)
	}
}

func send(msg string) {
	goui.RequestJSService(goui.JSServiceOptions{
		Url: "chat/" + msg,
	})
}
