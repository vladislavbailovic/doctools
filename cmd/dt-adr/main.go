package main

import (
	"doctools/pkg/cli"
	_ "embed"
	"os"
)

//go:embed resources/help.txt
var help string

func showHelp() {
	cli.Say(help)
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
		case "edit":
			editExisting(os.Args[2:])
		case "list", "ls":
			listAdrs(os.Args[2:])
		default:
			if len(os.Args) > 2 {
				changeAdrStatus(os.Args[1:])
			} else {
				showHelp()
			}
		}
	}
}
