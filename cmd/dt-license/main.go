package main

import (
	"doctools/pkg/cli"
	"embed"
	_ "embed"
	"fmt"
	"os"
)

//go:embed resources/help.txt
var help string

//go:embed resources/license_*.txt
var licenseFS embed.FS

func showHelp() {
	cli.Say(help)
}

func main() {
	if !cli.HasSubcommand() {
		showHelp()
		return
	} else {
		switch cli.Subcommand() {
		case "-h", "--help", "help":
			showHelp()
			return
		default:
			tpl := cli.Subcommand()
			switch tpl {
			case "gpl":
				tpl = "gpl2"
			case "bsd":
				tpl = "bsd3"
			case "lgpl":
				tpl = "lgpl3"
			}
			license, err := licenseFS.ReadFile(fmt.Sprintf("resources/license_%s.txt", tpl))
			if err != nil {
				cli.Cry("%v", err)
				cli.Nit("Falling back to wtfpl")
				license, _ = licenseFS.ReadFile("resources/license_wtf.txt")
			}
			if cli.HasFlag("-p") || cli.HasFlag("--print") || cli.HasFlag("print") {
				cli.Say(string(license))
				return
			}

			_, err = os.Stat("LICENSE.txt")
			if err == nil && !cli.HasFlag("-f") && !cli.HasFlag("--force") {
				cli.Cry("LICENSE.txt already exists")
				cli.Say("You can forcefully overwrite it, though (-f/--force)")
				return
			}
			if err := os.WriteFile("LICENSE.txt", license, 0622); err != nil {
				cli.Cry("%v", err)
			}
		}
	}
}
