package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/cli"
	"doctools/pkg/config"
	"strings"
)

func createNewAdr(args []string) {
	cfg, err := config.Load()
	if err != nil {
		cli.Cry("%v", err)
		return
	}
	create(cfg, strings.Join(args, " "))
}

func create(cfg config.Configuration, title string) {
	repo, err := adr.GetRepo(cfg)
	if err != nil {
		cli.Cry("error getting adr repo: %v", err)
		return
	}
	next, err := repo.NextID()
	if err != nil {
		cli.Cry("error getting next ID: %v", err)
		return
	}
	data := adr.Data{
		Number: next,
		Title:  title,
		Status: []adr.Status{adr.Status{Kind: adr.Drafted, Date: "today"}},
	}
	if err := adr.Save(data, repo); err != nil {
		cli.Cry("error saving adr to repo: %v", err)
		return
	}
	cli.Say("Created ADR: %d", next)

	openForEditing(data, repo)
}
