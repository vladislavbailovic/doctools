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

func main() {
	if !cli.HasSubcommand() {
		showHelp()
	} else {
		root, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			cli.Cry("%v", err)
			return
		}
		switch cli.Subcommand() {
		case "-h", "--help", "help":
			showHelp()
		case "adr":
			cmd := cli.Run(filepath.Join(root, "./dt-adr"), cli.SubcommandArgs()...)
			if err := cmd.Wait(); err != nil {
				cli.Cry("error executing: %v", err)
				return
			}
		case "rdme", "readme":
			cmd := cli.Run(filepath.Join(root, "./dt-rdme"), cli.SubcommandArgs()...)
			if err := cmd.Wait(); err != nil {
				cli.Cry("error executing: %v", err)
				return
			}
		default:
			showHelp()
		}
	}
}
