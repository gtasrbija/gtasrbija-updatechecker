package main

// #cgo CFLAGS: -g -Wall
import "C"

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nanitefactory/winmb"
)

const versionFileURL = "https://raw.githubusercontent.com/MikMik1011/gtasrbija-updatechecker/master/info/version.json"

func fetchVersion() string {
	res, err := http.Get(versionFileURL)
	if err != nil {
		winmb.MessageBoxPlain("GTA Srbija Update Checker", "Greska prilikom dobijanja informacije o najnovijoj verziji!")
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		winmb.MessageBoxPlain("GTA Srbija Update Checker", "Greska prilikom parsovanja informacije o najnovijoj verziji!")
		os.Exit(1)
	}

	return string(body)

}

//export MainThread
func MainThread() {
	winmb.MessageBoxPlain("GTA Srbija Update Checker", fetchVersion())

}

func main() {
	// nothing really. xD
}
