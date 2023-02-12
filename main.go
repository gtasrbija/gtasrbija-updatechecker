package main

// #cgo CFLAGS: -g -Wall
import "C"

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nanitefactory/winmb"
	"golang.org/x/mod/semver"
)

const currentVersion = "v0.0.1"

const versionFileURL = "https://raw.githubusercontent.com/MikMik1011/gtasrbija-updatechecker/master/info/version.txt"

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
	newestVersion := fetchVersion()

	diff := semver.Compare(currentVersion, newestVersion)

	switch diff {
	case -1:
		winmb.MessageBoxPlain("GTA Srbija Update Checker", "Nova verzija je dostupna!")
	case 1:
		winmb.MessageBoxPlain("GTA Srbija Update Checker", "Koristite noviju verziju moda nego sto je izasla!")
	}

}

func main() {
	// nothing really. xD
}
