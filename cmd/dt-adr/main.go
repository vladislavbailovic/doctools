package main

import (
	"doctools/pkg/cli"
	_ "embed"
)

//go:embed resources/help.txt
var help string

func showHelp() {
	cli.Say(help)
}

func main() {
	if !cli.HasSubcommand() {
		showHelp()
	} else {
		switch cli.Subcommand() {
		case "-h", "--help", "help":
			showHelp()
		case "init":
			initializeRepo()
		case "new", "draft", "create":
			createNewAdr(cli.SubcommandArgs())
		case "edit":
			editExisting(cli.SubcommandArgs())
		case "list", "ls":
			listAdrs(cli.SubcommandArgs())
		default:
			if cli.HasSubcommandArgs() {
				changeAdrStatus(cli.ArgsFrom(1))
			} else {
				showHelp()
			}
		}
	}
}
