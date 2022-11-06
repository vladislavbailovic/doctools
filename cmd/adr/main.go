package main

import (
	_ "embed"
	"os"

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

func showHelp() {
	dbg.Debug("showing help")
	dbg.Debug("%v", help)
}

func createNewAdr(args []string) {
	dbg.Debug("creating new")
	dbg.Debug("%#v", args)
	return

	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	create(cfg, "New ADR for testing")
}

func changeAdrStatus(args []string) {
	dbg.Debug("changing status")
	dbg.Debug("%#v", args)
	return

	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	changeStatus(cfg, 1, adr.Proposed)
}

func initializeRepo() {
	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	initialize(cfg)
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
	} else {
		switch os.Args[1] {
		case "-h", "--help", "help":
			showHelp()
		case "init":
			initializeRepo()
		case "new", "draft", "create":
			createNewAdr(os.Args[2:])
		default:
			if len(os.Args) > 2 {
				changeAdrStatus(os.Args[2:])
			} else {
				showHelp()
			}
		}
	}
}
