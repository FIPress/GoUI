//+build windows

package goui

//todo: c implementation of linux

import "C"

type window struct {
}

func (w *window) create(settings Settings) {
	//C.Create((*C.WindowSettings)(unsafe.Pointer(settings)))
	cs := convertSettings(settings)
	C.Create(cs)
}

func (w *window) activate() {

}

func (w *window) invokeJS(js string) {
	cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	C.InvokeJS(cJs)
}

func (w *window) exit() {
	C.Exit()
}
