package main

import (
	"fmt"
	"path/filepath"
)

type readme struct {
	Name        string
	Description string
	Sections    map[string][]string
}

func (x *readme) addSection(s section) {
	if x.Sections == nil {
		x.Sections = make(map[string][]string)
	}
	x.Sections[s.Name] = append(x.Sections[s.Name], s.Content)
}

type section struct {
	Name    string
	Content string
}

func newBuildSection(content string) section {
	return section{Name: "Building", Content: content}
}

func newRunSection(content string) section {
	return section{Name: "Running", Content: content}
}

func newTestSection(content string) section {
	return section{Name: "Testing", Content: content}
}

func newDockerSections(dockerfile string) []section {
	name := filepath.Base(filepath.Dir(dockerfile))
	if "" == name || "." == name {
		name = "latest"
		dockerfile = "."
	}
	return []section{
		newBuildSection(fmt.Sprintf("docker build %s -t %s", dockerfile, name)),
		newRunSection("docker run latest"),
	}
}

func newPhpunitSection(testfile string) section {
	return newTestSection(fmt.Sprintf("phpunit -c %s", testfile))
}
