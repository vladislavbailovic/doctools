package main

import (
	"doctools/pkg/cli"
	"strings"
)

func main() {
	cli.Say(
		"%#v",
		getLogLines(),
	)
}

func getLogLines(extraParams ...string) []string {
	params := []string{"log", "--oneline"}
	params = append(params, extraParams...)
	out := cli.CaptureOutput("git", params...)
	return strings.Split(strings.TrimSpace(out), "\n")
}
