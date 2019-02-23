package goui

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Context struct {
	Data            string `json:"data"`
	SuccessCallback string `json:"success"`
	ErrorCallback   string `json:"error"`
	params          map[string]string
}

func (ctx *Context) GetParam(name string) string {
	return ctx.params[name]
}

func (ctx *Context) GetBoolParam(name string) (b bool, err error) {
	str := ctx.GetParam(name)
	b, err = strconv.ParseBool(str)
	if err != nil {
		Log("convert data to bool failed:", err)
	}
	return
}

func (ctx *Context) GetIntParam(name string) (i int, err error) {
	str := ctx.GetParam(name)
	i, err = strconv.Atoi(str)
	if err != nil {
		Log("convert data to int failed:", err)
	}
	return
}

func (ctx *Context) GetFloatParam(name string) (f float64, err error) {
	str := ctx.GetParam(name)
	f, err = strconv.ParseFloat(str, 32)
	if err != nil {
		Log("convert data to float failed:", err)
	}
	return
}

func (ctx *Context) GetEntity(v interface{}) (err error) {
	err = json.Unmarshal([]byte(ctx.Data), v)
	if err != nil {
		Log("get entity failed:", err)
	}
	return
}

func (ctx *Context) Success(feedback interface{}) {
	if ctx.SuccessCallback != "" {
		if feedback == nil {
			w.invokeJS(ctx.SuccessCallback + "()")
		} else {
			w.invokeJS(fmt.Sprintf("%s('%v')", ctx.SuccessCallback, feedback))
		}
	}
}

func (ctx *Context) Error(err interface{}) {
	if ctx.ErrorCallback != "" {
		if err == nil {
			w.invokeJS(ctx.ErrorCallback + "()")
		} else {
			w.invokeJS(fmt.Sprintf("%s('%v')", ctx.ErrorCallback, err))
		}
	}
}

func (ctx *Context) Ok() {
	ctx.Success(nil)
}

func (ctx *Context) Fail() {
	ctx.Error(nil)
}
