package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"strconv"
)

func changeAdrStatus(args []string) {
	status, err := adr.StatusTypeFromString(args[0])
	if err != nil {
		dbg.Error("%v", err)
		showHelp()
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		dbg.Error("%v", err)
		showHelp()
		return
	}

	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}

	changeStatus(cfg, uint(num), status)
}

func changeStatus(cfg config.Configuration, number uint, status adr.StatusType) {
	repo, err := adr.GetRepo(cfg)
	if err != nil {
		dbg.Error("error getting adr repo: %v", err)
		return
	}
	data, err := repo.GetByNumber(number)
	if err != nil {
		dbg.Error("error getting adr data: %v", err)
		return
	}

	updated := data.UpdateStatus(status)
	if err := adr.Save(updated, repo); err != nil {
		dbg.Error("error saving updated adr: %v", err)
		return
	}
	dbg.Debug("Updated ADR status: %#v", updated.Status)
}
