package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"strings"
)

func createNewAdr(args []string) {
	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	create(cfg, strings.Join(args, " "))
}

func create(cfg config.Configuration, title string) {
	repo, err := adr.GetRepo(cfg)
	if err != nil {
		dbg.Error("error getting adr repo: %v", err)
		return
	}
	next, err := repo.NextID()
	if err != nil {
		dbg.Error("error getting next ID: %v", err)
		return
	}
	data := adr.Data{
		Number: next,
		Title:  title,
		Status: []adr.Status{adr.Status{Kind: adr.Drafted, Date: "today"}},
	}
	if err := adr.Save(data, repo); err != nil {
		dbg.Error("error saving adr to repo: %v", err)
		return
	}
	dbg.Debug("Created ADR: %d", next)

	openForEditing(data, repo)
}
