package main

import "doctools/pkg/cli"

func main() {
	if !cli.HasSubcommand() {
		cli.Say("HALP!")
	} else {
		switch cli.Subcommand() {
		case "adr":
			cmd := cli.Run("./dt-adr", cli.SubcommandArgs()...)
			if err := cmd.Wait(); err != nil {
				cli.Cry("error executing: %v", err)
				return
			}
		default:
			cli.Say("HALP!")
		}
	}
}
