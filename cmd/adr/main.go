package main

import (
	_ "embed"
	"fmt"

	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"doctools/pkg/storage"
)

//go:embed resources/help.txt
var help string

func initialize(cfg config.Configuration) {
	if err := config.InitializeProject(cfg); err != nil {
		dbg.Error("error initializing project config: %v", err)
		return
	}
	if err := storage.InitializeProject(cfg); err != nil {
		dbg.Error("error initializing project storage: %v", err)
		return
	}

	repo, err := adr.GetRepo(cfg)
	if err != nil {
		if err := repo.Initialize(cfg); err != nil {
			dbg.Error("error initializing ADR storage: %v", err)
			return
		}
	}

	dbg.Debug("Created subdirectory for ADRs")
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
}

func main() {
	fmt.Printf("Help, then [%v]", help)

	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	initialize(cfg)

	create(cfg, "New ADR for testing")
	/*
		data := adr.Data{
			Title: "Use ADRs",
			Status: []adr.Status{
				adr.Status{Kind: adr.Drafted, Date: "2022-11-05"},
				adr.Status{Kind: adr.Proposed, Date: "2022-11-05"},
			},
			Context:      "Keeping track of *why* something was done gets error-prone over time.",
			Decision:     "We're gonna start using ADRs to document major decisions",
			Consequences: "Start doctools utility development to facilitate this",
		}
		repo, err := adr.GetRepo(cfg)
		if err != nil {
			dbg.Error("error getting adr repo: %v", err)
			return
		}
		if err := adr.Save(data, repo); err != nil {
			dbg.Error("error saving adr to repo: %v", err)
			return
		}
		dbg.Debug("Saved ADR")
	*/
}
