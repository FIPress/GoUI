package goui

//todo: c implementation of linux

import "C"

type windowsWorker struct {
	Settings
}

func (w *windowsWorker) create() {
	/*C.InitApp()

	cTitle := C.CString(w.Title)
	defer C.free(unsafe.Pointer(cTitle))

	cUrl := C.CString(w.Url)
	defer C.free(unsafe.Pointer(cUrl))

	println("url:",w.Url)

	C.CreateWindow(cTitle,cUrl,C.int(w.Width),C.int(w.Height))

	C.ActivateApp()

	for C.Loop() == 0 {}*/
}

func (w *windowsWorker) invokeJS(js string) {
	/*cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	C.InvokJS(cJs)*/
}

func (w *windowsWorker) exit() {
	/*println("close cocoa window")
	C.Exit()*/
}
