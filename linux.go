package goui

//todo: c implementation of linux

import "C"

type linuxWorker struct {
	Settings
}

func (w *linuxWorker) create() {
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

func (w *linuxWorker) invokeJS(js string) {
	/*cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	C.InvokJS(cJs)*/
}

func (w *linuxWorker) exit() {
	/*println("close linux window")
	C.Exit()*/
}
