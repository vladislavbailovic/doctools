package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/output"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed resources/list.txt
var templateSource string
var tpl = template.Must(
	template.New("List").Parse(templateSource),
)

func listAdrs(args []string) {
	cfg, err := config.Load()
	if err != nil {
		output.Cry("%v", err)
		return
	}

	repo, err := adr.GetRepo(cfg)
	if err != nil {
		output.Cry("error getting adr repo: %v", err)
		return
	}

	var list []adr.Data

	all, err := repo.ListAll()
	if err != nil {
		output.Cry("error listing ADRs: %v", err)
		return
	}

	if len(args) == 0 {
		list = all
	} else {
		status, err := adr.StatusTypeFromString(args[0])
		if err != nil {
			output.Cry("error filtering ADRs: %v", err)
			return
		}
		for _, data := range all {
			if len(data.Status) == 0 || data.Status[len(data.Status)-1].Kind != status {
				continue
			}
			list = append(list, data)
		}
	}

	if len(list) == 0 {
		output.Say("No ADRs")
		return
	}

	buffer := new(strings.Builder)
	tpl.Execute(buffer, list)
	output.Say(strings.TrimSpace(buffer.String()))
}
