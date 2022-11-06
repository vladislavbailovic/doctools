package main

import (
	_ "embed"
	"os"

	"doctools/pkg/dbg"
)

//go:embed resources/help.txt
var help string

func showHelp() {
	dbg.Debug("showing help")
	dbg.Debug("%v", help)
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
	} else {
		switch os.Args[1] {
		case "-h", "--help", "help":
			showHelp()
		case "init":
			initializeRepo()
		case "new", "draft", "create":
			createNewAdr(os.Args[2:])
		default:
			if len(os.Args) > 2 {
				changeAdrStatus(os.Args[2:])
			} else {
				showHelp()
			}
		}
	}
}
