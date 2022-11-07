package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/cli"
	"doctools/pkg/config"
	"strconv"
)

func changeAdrStatus(args []string) {
	status, err := adr.StatusTypeFromString(args[0])
	if err != nil {
		cli.Cry("%v", err)
		showHelp()
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		cli.Cry("%v", err)
		showHelp()
		return
	}

	cfg, err := config.Load()
	if err != nil {
		cli.Cry("%v", err)
		return
	}

	changeStatus(cfg, uint(num), status)
}

func changeStatus(cfg config.Configuration, number uint, status adr.StatusType) {
	repo, err := adr.GetRepo(cfg)
	if err != nil {
		cli.Cry("error getting adr repo: %v", err)
		return
	}
	data, err := repo.GetByNumber(number)
	if err != nil {
		cli.Cry("error getting adr data: %v", err)
		return
	}

	updated := data.UpdateStatus(status)
	if err := adr.Save(updated, repo); err != nil {
		cli.Cry("error saving updated adr: %v", err)
		return
	}
	cli.Nit("Updated ADR status: %#v", updated.Status)
}
