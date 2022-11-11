package main

import (
	"doctools/pkg/cli"
	"doctools/pkg/markdown"
	"os"
	"strings"
)

func main() {
	// changeset := getChangesets()
	changeset := fromChangelog()
	renderChangeset(changeset)
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

func fromChangelog() []changeset {
	result := []changeset{}

	raw, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		return result
	}
	lines := strings.Split(string(raw), "\n")
	md := markdown.NewMarkdownFromLines(lines)

	pos := md.FindHeader(markdown.HeaderAny)
	for pos >= 0 {
		next := md.FindHeaderAfter(pos, markdown.HeaderAny)
		current := lines[pos]

		content := []string{}
		if next > 0 {
			content = lines[pos+1 : next]
		} else {
			content = lines[pos+1:]
		}

		set := changeset{
			name: markdown.GetHeaderText(current),
		}
		for _, c := range content {
			if len(c) == 0 {
				continue
			}
			set.changes = append(set.changes, markdown.Delistify(c))
		}
		result = append(result, set)

		pos = next
	}

	return result
}
