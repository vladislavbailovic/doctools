package main

import (
	"doctools/pkg/cli"
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed resources/help.txt
var help string

func showHelp() {
	cli.Say(help)
}

var cmds map[string]string = map[string]string{
	"adr":       "dt-adr",
	"rdme":      "dt-rdme",
	"readme":    "dt-rdme",
	"lcs":       "dt-license",
	"license":   "dt-license",
	"chglg":     "dt-chglg",
	"changelog": "dt-chglg",
}

func main() {
	if !cli.HasSubcommand() {
		showHelp()
	} else {
		root, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			cli.Cry("%v", err)
			return
		}
		subcommand := cli.Subcommand()
		switch subcommand {
		case "-h", "--help", "help":
			showHelp()
		default:
			for alias, command := range cmds {
				if subcommand != alias {
					continue
				}
				cmd := cli.Run(filepath.Join(root, command), cli.SubcommandArgs()...)
				if err := cmd.Wait(); err != nil {
					cli.Cry("error executing: %v", err)
					return
				}
			}
			showHelp() // If we got here, no good
			return
		}
	}
}
