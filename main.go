package main

// #cgo CFLAGS: -g -Wall
import "C"

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/browser"
	"github.com/sqweek/dialog"
	"golang.org/x/mod/semver"
)

const currentVersion = "v1.0.0"

const versionFileURL = "https://raw.githubusercontent.com/gtasrbija/gtasrbija-updatechecker/master/info/version.txt"
const downloadURL = "https://www.gtasrbija.net/download"

const dialTitle = "GTA Srbija Update Checker"
const reqErrText = "Greska prilikom dobijanja informacije o najnovijoj verziji!"
const readErrText = "Greska prilikom obrade informacije o najnovijoj verziji!"
const newVerText = "Nova verzija moda je dostupna! \n\nTrenutna verzija: %s \nNova verzija: %s \n\nDa li zelite da preuzmete novu verziju?"
const unrelText = "Koristite noviju, neobjavljenu verziju! \nImajte na umu da je ZABRANJENO leakovanje novih stvari bez odobrenja autora!"

func fetchVersion() string {
	res, err := http.Get(versionFileURL)
	if err != nil {
		dialog.Message(reqErrText).Title(dialTitle).Error()
		return ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		dialog.Message(readErrText).Title(dialTitle).Error()
		return ""
	}

	return string(body)

}

//export MainThread
func MainThread() {
	newestVersion := fetchVersion()
	if !semver.IsValid(currentVersion) || !semver.IsValid(newestVersion) {
		return
	}

	diff := semver.Compare(currentVersion, newestVersion)

	switch diff {
	case -1:
		if dialog.Message(fmt.Sprintf(newVerText, currentVersion, newestVersion)).Title(dialTitle).YesNo() {
			browser.OpenURL(downloadURL)
		}

	case 1:
		dialog.Message(unrelText).Title(dialTitle).Info()
	}

}

func main() {
	// nothing really. xD
}
