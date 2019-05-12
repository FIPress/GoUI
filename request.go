package goui

import "C"

import (
	"encoding/json"
)

type Request struct {
	Url        string `json:"url"`
	IsCallback bool   `json:"isCallback"`
	Context    `json:",inline"`
}

var cbMap = make(map[string]func(string))

//export handleClientReq
func handleClientReq(msg *C.char) {
	message := C.GoString(msg)
	Log("ClientHandler:", message)

	req := new(Request)
	err := json.Unmarshal([]byte(message), req)
	if err != nil {
		Log("unmarshal error:", err)
		return
	}

	if req.IsCallback {
		f := cbMap[req.Url]
		if f != nil {
			f(req.Data)
			delete(cbMap, req.Url)
		}
	} else {
		ctx := &req.Context
		handler, params := dispatch(req.Url)
		if handler != nil {
			ctx.params = params
			handler(ctx)
		} else {
			notFound(ctx)
		}
	}

}

func notFound(ctx *Context) {
	ctx.Error("Function not found ")
}
