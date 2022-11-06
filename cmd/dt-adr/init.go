package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"doctools/pkg/storage"
)

func initializeRepo() {
	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	initialize(cfg)
}

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
