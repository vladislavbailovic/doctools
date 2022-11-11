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
	if !cli.HasSubcommand() {
		showHelp()
	} else {
		switch cli.Subcommand() {
		case "-h", "--help", "help":
			showHelp()
		case "new", "init":
			_, err := os.Stat("CHANGELOG.md")
			if err == nil && !cli.HasFlag("-f") && !cli.HasFlag("--force") {
				cli.Cry("CHANGELOG.md already exists")
				cli.Say("You can forcefully overwrite it, though (-f/--force)")
				return
			}

			if err := initChangelog(); err != nil {
				cli.Cry("%v", err)
			}
		case "update", "toc":
			_, err := os.Stat("CHANGELOG.md")
			if err != nil {
				if err := initChangelog(); err != nil {
					cli.Cry("%v", err)
				}
				return
			}
			if err := updateChangelog(); err != nil {
				cli.Cry("%v", err)
			}
		default:
			showHelp()
		}
	}
}

func initChangelog() error {
	wip := fromRepo()
	return writeChangelog(wip)
}

func updateChangelog() error {
	known := fromFile("CHANGELOG.md")
	wip := fromRepo()
	final := known.updateFrom(wip)
	return writeChangelog(final)
}

func writeChangelog(log changelog) error {
	return os.WriteFile("CHANGELOG.md", []byte(log.String()), 0622)
}
