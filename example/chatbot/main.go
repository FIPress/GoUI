package main

import (
	"github.com/fipress/GoUI"
	"strings"
	"time"
)

const (
	greetings = "Hello, I'm your pal B. How may I help you today?"
	boast     = "You can ask me anything. There is nothing I don't know."
	truth     = "Uh-oh~~~  you got me! I'm just a dummy robot."
	knock     = "Knock knock..."
	please    = "Say something please..."
)

var timer *time.Timer

func main() {
	goui.Service("chat/:msg", chatService)
	timer = time.NewTimer(time.Minute)
	resetTimer()
	goui.Create(goui.Settings{Title: "Chatbot", Top: 30, Left: 100, Width: 300, Height: 440})
}

func resetTimer() {
	timer.Reset(time.Minute)

	go func() {
		select {
		case <-timer.C:
			println("timeout")
			goui.RequestJSService(goui.JSServiceOptions{
				Url: "chat/" + knock,
			}, false)
			resetTimer()
		}
	}()
}

func chatService(ctx *goui.Context) {
	msg := ctx.GetParam("msg")
	resetTimer()
	msg = strings.ToLower(msg)
	switch {
	case msg == "":
		ctx.Success(please)
	case msg == "ready":
		ctx.Success(greetings)
		goui.RequestJSService(goui.JSServiceOptions{
			Url: "chat/" + greetings,
		}, true)
	case strings.HasSuffix(msg, "?"):
		ctx.Success(truth)
	default:
		ctx.Success(boast)
	}
}

/*
func receive(msg string) {
	resetTimer()
	msg = strings.ToLower(msg)
	switch {
	case msg == "":
		send(please)
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
}*/
