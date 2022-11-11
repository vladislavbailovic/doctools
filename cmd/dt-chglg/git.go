package main

import (
	"doctools/pkg/cli"
	"strings"
)

type tag struct {
	commit
	name string
}

func getTags() []tag {
	tags := []tag{}

	for _, line := range getTagNames() {
		tagName := strings.TrimSpace(line)
		if "" == tagName {
			continue
		}
		nfo := getCommits(tagName, "-n", "1")
		if len(nfo) != 1 {
			cli.Nit("parsing tag %s, wrong result %#v", tagName, nfo)
			continue
		}
		tag := tag{
			commit: nfo[0],
			name:   tagName,
		}
		tags = append(tags, tag)
	}
	return tags
}

func getTagNames() []string {
	raw := strings.TrimSpace(cli.CaptureOutput("git", "tag"))
	if "" == raw {
		return []string{}
	}
	return strings.Split(raw, "\n")
}

func getTagDate(tag string) string {
	raw := strings.TrimSpace(cli.CaptureOutput("git", "log", "-1", "--pretty='%cI'", tag))
	return strings.Replace(raw, "'", "", -1)
}

type commit struct {
	hash  string
	title string
}

func getCommits(extraParams ...string) []commit {
	result := []commit{}
	lines := getLogLines(extraParams...)
	if len(lines) == 0 {
		return result
	}
	for _, line := range lines {
		cmt := commitFromLogLine(line)
		if cmt == (commit{}) {
			continue
		}
		result = append(result, cmt)
	}
	return result
}

func getCommitsBetween(last string, first ...string) []commit {
	result := []commit{}

	params := []string{}
	if len(first) == 1 {
		params = append(params, first[0])
	} else if len(first) > 1 {
		cli.Nit("expected zero or more optional hashes, got %#v", first)
		return result
	}
	params = append(params, last)

	result = getCommits(strings.Join(params, ".."))
	return result
}

func commitFromLogLine(line string) commit {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) != 2 {
		return commit{}
	}
	title := strings.TrimSpace(parts[1])
	if strings.HasPrefix(title, "(tag: ") {
		tagEnd := strings.Index(title, ")")
		if tagEnd > 0 && tagEnd < len(title) {
			title = strings.TrimSpace(title[tagEnd:])
		}
	}
	return commit{
		hash:  strings.TrimSpace(parts[0]),
		title: strings.TrimSpace(parts[1]),
	}
}

func getLogLines(extraParams ...string) []string {
	params := []string{"log", "--oneline"}
	params = append(params, extraParams...)
	out := cli.CaptureOutput("git", params...)
	if "" == out {
		return []string{}
	}
	return strings.Split(strings.TrimSpace(out), "\n")
}

func firstCommitDescriptor() string {
	return strings.TrimSpace(cli.CaptureOutput("git", "rev-list", "--max-parents=0", "HEAD"))
}

func lastCommitDescriptor() string {
	return "HEAD"
}
