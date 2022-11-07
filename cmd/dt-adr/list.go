package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
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
		dbg.Error("%v", err)
		return
	}

	repo, err := adr.GetRepo(cfg)
	if err != nil {
		dbg.Error("error getting adr repo: %v", err)
		return
	}

	var list []adr.Data

	all, err := repo.ListAll()
	if err != nil {
		dbg.Error("error listing ADRs: %v", err)
		return
	} else {
		list = all
	}

	if len(list) == 0 {
		dbg.Debug("No ADRs")
		return
	}

	buffer := new(strings.Builder)
	tpl.Execute(buffer, list)
	dbg.Debug("\n" + strings.TrimSpace(buffer.String()))
}
