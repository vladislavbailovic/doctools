package main

import (
	"doctools/pkg/cli"
)

func main() {
	known := fromFile("CHANGELOG.md")
	wip := fromRepo()
	changeset := known.updateFrom(wip)
	renderChangeset(changeset.changes)
}

func renderChangeset(all []changeset) {
	for _, set := range all {
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

func diffChangesets(wip, known []changeset) []changeset {
	if len(known) > len(wip) {
		return []changeset{}
	}
	diff := len(wip) - len(known)
	if diff == 0 {
		return known
	}

	return wip[:diff]
}
