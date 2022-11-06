package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"strconv"
	"strings"
)

func changeAdrStatus(args []string) {
	var status adr.StatusType
	switch strings.ToLower(args[0]) {
	case "draft":
		status = adr.Drafted
	case "propose":
		status = adr.Proposed
	case "accept":
		status = adr.Accepted
	case "reject":
		status = adr.Rejected
	case "supersede":
		status = adr.Superseded
	default:
		dbg.Error("unknown new status: %v", args[0])
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
