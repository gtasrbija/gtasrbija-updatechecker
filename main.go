package main

// #cgo CFLAGS: -g -Wall
import "C"

import (
	"github.com/nanitefactory/winmb"
)

//export MainThread
func MainThread() {
	winmb.MessageBoxPlain("GTA Srbija Update Checker", "Proba proba 1312")
}

func main() {
	// nothing really. xD
}
