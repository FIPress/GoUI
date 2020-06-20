package goui

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Context struct {
	Data            string `json:"data"`
	SuccessCallback string `json:"success"`
	ErrorCallback   string `json:"error"`
	params          map[string]string
}

// GetParam get a string parameter from the url
func (ctx *Context) GetParam(name string) (val string) {
	val, _ = url.PathUnescape(ctx.params[name])
	return
}

// GetBoolParam get a bool parameter from the url
func (ctx *Context) GetBoolParam(name string) (b bool, err error) {
	str := ctx.GetParam(name)
	b, err = strconv.ParseBool(str)
	if err != nil {
		Log("convert data to bool failed:", err)
	}
	return
}

// GetIntParam get a int parameter from the url
func (ctx *Context) GetIntParam(name string) (i int, err error) {
	str := ctx.GetParam(name)
	i, err = strconv.Atoi(str)
	if err != nil {
		Log("convert data to int failed:", err)
	}
	return
}

func (ctx *Context) GetIntParamOr(name string, defaultVal int) (i int) {
	str := ctx.GetParam(name)
	var err error
	i, err = strconv.Atoi(str)
	if err != nil {
		i = defaultVal
	}
	return

}

// GetFloatParam get a float parameter from the url
func (ctx *Context) GetFloatParam(name string) (f float64, err error) {
	str := ctx.GetParam(name)
	f, err = strconv.ParseFloat(str, 32)
	if err != nil {
		Log("convert data to float failed:", err)
	}
	return
}

// GetParam get an entity from the requested data
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
			invokeJS(ctx.SuccessCallback+"()", true)
		} else {
			invokeJS(fmt.Sprintf("%s(\"%v\")", ctx.SuccessCallback, feedback), true)
		}
	}
}

func (ctx *Context) Error(err interface{}) {
	if ctx.ErrorCallback != "" {
		if err == nil {
			invokeJS(ctx.ErrorCallback+"()", true)
		} else {
			invokeJS(fmt.Sprintf("%s('%v')", ctx.ErrorCallback, err), true)
		}
	}
}

func (ctx *Context) Ok() {
	ctx.Success(nil)
}

func (ctx *Context) Fail() {
	ctx.Error(nil)
}
