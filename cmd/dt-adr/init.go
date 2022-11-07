package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/cli"
	"doctools/pkg/config"
	"doctools/pkg/storage"
)

func initializeRepo() {
	cfg, err := config.Load()
	if err != nil {
		cli.Cry("%v", err)
		return
	}
	initialize(cfg)
}

func initialize(cfg config.Configuration) {
	if err := config.InitializeProject(cfg); err != nil {
		cli.Cry("error initializing project config: %v", err)
		return
	}
	if err := storage.InitializeProject(cfg); err != nil {
		cli.Cry("error initializing project storage: %v", err)
		return
	}

	repo, err := adr.GetRepo(cfg)
	if err != nil {
		if err := repo.Initialize(cfg); err != nil {
			cli.Cry("error initializing ADR storage: %v", err)
			return
		}
	}

	cli.Nit("Created subdirectory for ADRs")
}
