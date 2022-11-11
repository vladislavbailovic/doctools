package main

import (
	"doctools/pkg/cli"
)

func main() {
	for _, set := range getChangesets() {
		cli.Say("### %s", set.name)
		if set.hasChanges() {
			for _, c := range set.changes {
				cli.Say("\t- %s", c)
			}
		} else {
			cli.Say("\t* No changes *")
		}
		cli.Say("")
	}
}
