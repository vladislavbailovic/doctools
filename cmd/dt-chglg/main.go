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

func reverse(input []changeset) []changeset {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

type changeset struct {
	name    string
	changes []string
}

func (x changeset) hasChanges() bool {
	return len(x.changes) > 0
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
