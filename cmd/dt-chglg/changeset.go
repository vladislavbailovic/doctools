package main

import (
	"doctools/pkg/cli"
	"doctools/pkg/markdown"
	"strings"
	"time"
)

const dateFormat string = "2006-01-02"

type changeset struct {
	name    string
	date    time.Time
	changes []string
}

func (x changeset) hasChanges() bool {
	return len(x.changes) > 0
}

func (x changeset) String() string {
	result := make([]string, len(x.changes)+2, len(x.changes)+2)
	result[0] = strings.Join([]string{
		markdown.HeaderLevel3.String(),
		" ",
		x.name,
		" ",
		"(" + x.date.Format(dateFormat) + ")",
	}, "")
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
	date, err := time.Parse(time.RFC3339, getTagDate(now))
	if err != nil {
		cli.Nit("error for date string [%s]: %v", getTagDate(now), err)
	}
	return changeset{
		name:    name,
		date:    date,
		changes: getChangesBetween(now, since),
	}
}

func getWIPChangeset() changeset {
	date, err := time.Parse(time.RFC3339, getTagDate("HEAD"))
	if err != nil {
		cli.Nit("error for date string [%s]: %v", getTagDate("HEAD"), err)
	}
	return changeset{
		name:    "WIP",
		date:    date,
		changes: getWIPChanges(),
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
	name, date := parseChangesetName(name)
	return changeset{
		name:    name,
		date:    date,
		changes: result,
	}
}

func parseChangesetName(name string) (string, time.Time) {
	parts := strings.Split(name, "(")
	if len(parts) == 2 {
		name = strings.TrimSpace(parts[0])
		date, err := time.Parse(dateFormat, parts[1][:len(parts[1])-1])
		if err != nil {
			cli.Nit("error for date string [%s]: %v", parts[1][:len(parts[1])-1], err)
		}
		return name, date
	}
	return name, time.Time{}
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
