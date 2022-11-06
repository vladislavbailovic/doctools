package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"doctools/pkg/project"
)

//go:embed resources/help.txt
var help string

const (
	AdrDirectory string = "adr"
)

func initialize(cfg config.Configuration) {
	if err := project.InitializeProjectDirectory(cfg); err != nil {
		dbg.Error("error initializing project: %v", err)
		return
	}

	docsDir, err := project.GetDocsDirectory(cfg)
	if err != nil {
		dbg.Error("error getting doc dir: %v", err)
		return
	}

	adrDir := filepath.Join(docsDir, AdrDirectory)
	if nfo, err := os.Stat(adrDir); err != nil {
		if err := os.Mkdir(adrDir, 0722); err != nil {
			dbg.Error("error creating ADR subdir: %v", err)
			return
		}
	} else if nfo.IsDir() {
		dbg.Debug("yay")
		return
	}

	if _, err := os.Stat(adrDir); err != nil {
		dbg.Error("ADR subdir not created (%s): %v", adrDir, err)
	}

	dbg.Debug("Created subdirectory for ADRs: %s", adrDir)
}

func main() {
	fmt.Printf("Help, then [%v]", help)

	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	initialize(cfg)

	adr := adr.Data{
		Title: "Use ADRs",
		Status: []adr.Status{
			adr.Status{Kind: adr.Drafted, Date: "2022-11-05"},
			adr.Status{Kind: adr.Proposed, Date: "2022-11-05"},
		},
		Context:      "Keeping track of *why* something was done gets error-prone over time.",
		Decision:     "We're gonna start using ADRs to document major decisions",
		Consequences: "Start doctools utility development to facilitate this",
	}
	dbg.Debug("\n\n%v\n\n", adr)
}
