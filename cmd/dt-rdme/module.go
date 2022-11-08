package main

import (
	"doctools/pkg/cli"
	"os"
	"strings"
)

type modulefile struct {
	module string
	cmds   []string
}

func projectModuleFile(p projectInfo) modulefile {
	var data modulefile
	raw, err := p.getFile("go.mod")
	if err != nil {
		cli.Nit("%v", err)
		return data
	}

	for _, line := range strings.Split(string(raw), "\n") {
		if strings.HasPrefix(line, "module ") {
			data.module = strings.TrimPrefix(line, "module ")
			break
		}
	}

	if p.hasDir("cmd") {
		if entries, err := os.ReadDir(p.getPath("cmd")); err == nil {
			for _, item := range entries {
				if !item.IsDir() {
					continue
				}
				data.cmds = append(data.cmds, item.Name())
			}
		}
	}

	return data
}
