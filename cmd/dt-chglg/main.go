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
		case "update", "sync":
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
		case "wip":
			cli.Say("%v", getWIPChangeset())
		case "show":
			args := cli.SubcommandArgs()
			if len(args) < 1 {
				showHelp()
				return
			}
			what := args[0]
			if what == "WIP" {
				cli.Say("%v", getWIPChangeset())
			} else {
				log := syncChangelog()
				cli.Say("%v", log.findChange(what))
			}
		default:
			showHelp()
		}
	}
}

func syncChangelog() changelog {
	known := fromFile("CHANGELOG.md")
	wip := fromRepo()
	return known.updateFrom(wip)
}

func initChangelog() error {
	wip := fromRepo()
	return writeChangelog(wip)
}

func updateChangelog() error {
	final := syncChangelog()
	return writeChangelog(final)
}

func writeChangelog(log changelog) error {
	return os.WriteFile("CHANGELOG.md", []byte(log.String()), 0622)
}
