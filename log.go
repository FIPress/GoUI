// +build !android

package goui

import "C"
import "fmt"

//export goLog
func goLog(msg *C.char) {
	go fmt.Println("go log: ", C.GoString(msg))
}

func Log(args ...interface{}) {
	go fmt.Print(args...)
}

func Logf(format string, args ...interface{}) {
	go fmt.Printf(format, args...)
}
