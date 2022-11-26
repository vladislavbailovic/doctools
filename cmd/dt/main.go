package main

import (
	"doctools/pkg/cli"
	_ "embed"
	"os"
	"path/filepath"
	"strings"
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
		root := ""
		if strings.HasPrefix(os.Args[0], os.TempDir()) {
			root = "./"
		}
		subcommand := cli.Subcommand()
		switch subcommand {
		case "-h", "--help", "help":
			showHelp()
		case "init":
			params := []string{"init"}
			if cli.HasFlag("-f") || cli.HasFlag("--force") {
				params = append(params, "--force")
			}
			executed := map[string]bool{}
			for _, command := range cmds {
				if _, ok := executed[command]; ok {
					continue
				}
				executed[command] = true
				cmd := cli.Run(filepath.Join(root, command), params...)
				if err := cmd.Wait(); err != nil {
					cli.Cry("error executing %s init: %v", command, err)
					continue
				}
			}
			return
		default:
			for alias, command := range cmds {
				if subcommand != alias {
					continue
				}
				cmd := cli.Run(filepath.Join(root, command), cli.SubcommandArgs()...)
				if err := cmd.Wait(); err != nil {
					cli.Cry("error executing: %v", err)
				}
				return
			}
			showHelp() // If we got here, no good
			return
		}
	}
}
