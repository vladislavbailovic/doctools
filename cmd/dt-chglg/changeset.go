package main

import (
	"doctools/pkg/markdown"
	"strings"
)

type changeset struct {
	name    string
	changes []string
}

func (x changeset) hasChanges() bool {
	return len(x.changes) > 0
}

func (x changeset) String() string {
	result := make([]string, len(x.changes)+2, len(x.changes)+2)
	result[0] = markdown.HeaderLevel3.String() + " " + x.name
	result[1] = ""

	for i, chg := range x.changes {
		result[i+2] = markdown.Listify(chg, 0)
	}

	return strings.Join(result, "\n")
}

func getChangesets() []changeset {
	result := []changeset{}
	tags := getTagNames()

	prev := firstCommitDescriptor()
	for _, tag := range tags {
		set := getChangeset(prev, tag)
		if set.hasChanges() {
			result = append(result, set)
		}
		prev = tag
	}
	set := getChangeset(prev, lastCommitDescriptor())
	if set.hasChanges() {
		result = append(result, set)
	}

	return reverse(result)
}

func getChangeset(since, now string) changeset {
	name := now
	if now == lastCommitDescriptor() {
		name = "WIP"
	}
	return changeset{
		name:    name,
		changes: getChangesBetween(now, since),
	}
}

func parseChangeset(name string, list []string) changeset {
	result := []string{}
	for _, item := range list {
		item = markdown.Delistify(item)
		if len(item) == 0 {
			continue
		}
		result = append(result, item)
	}
	return changeset{
		name:    name,
		changes: result,
	}
}

func getWIPChanges() []string {
	result := []string{}

	tags := getTagNames()
	if len(tags) < 1 {
		return result
	}

	tag := tags[len(tags)-1]
	return getChangesBetween(lastCommitDescriptor(), tag)
}

func getChangesBetween(earliest, oldest string) []string {
	result := []string{}
	for _, cmt := range getCommitsBetween(earliest, oldest) {
		result = append(result, cmt.title)
	}
	return result
}

func reverse(input []changeset) []changeset {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}
