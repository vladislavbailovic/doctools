package main

import (
	"doctools/pkg/markdown"
	"os"
	"strings"
)

type changelog struct {
	milestones map[string]bool
	changes    []changeset
}

func fromRepo() changelog {
	return fromChangesets(getChangesets())
}

func fromChangesets(changes []changeset) changelog {
	milestones := make(map[string]bool, len(changes))
	for _, item := range changes {
		milestones[item.name] = true
	}
	return changelog{changes: changes, milestones: milestones}
}

func fromFile(path string) changelog {
	changes := []changeset{}

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

		name := markdown.GetHeaderText(current)
		changes = append(changes, parseChangeset(name, content))

		pos = next
	}

	return fromChangesets(changes)
}

func (x changelog) updateFrom(wip changelog) changelog {
	result := []changeset{}
	for _, set := range wip.changes {
		if _, ok := x.milestones[set.name]; ok {
			break // break on first set with same milestone name
		}
		result = append(result, set)
	}
	for _, set := range x.changes {
		result = append(result, set)
	}
	return fromChangesets(result)
}
