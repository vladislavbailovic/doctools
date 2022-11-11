package main

import (
	"doctools/pkg/markdown"
	"os"
	"strings"
)

type changelog struct {
	changes []changeset
}

func fromFile(path string) changelog {
	result := []changeset{}

	raw, err := os.ReadFile(path)
	if err != nil {
		return changelog{}
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
			c = markdown.Delistify(c)
			if len(c) == 0 {
				continue
			}
			set.changes = append(set.changes, c)
		}
		result = append(result, set)

		pos = next
	}

	return changelog{changes: result}
}
