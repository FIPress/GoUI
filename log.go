package goui

import "C"
import "fmt"

//export goLog
func goLog(msg string) {
	fmt.Println(msg)
}

func Log(args ...interface{}) {
	fmt.Print(args...)
}

func Logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
