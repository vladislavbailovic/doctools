package main

import (
	"doctools/pkg"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"doctools/pkg/project"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed resources/help.txt
var help string

const (
	AdrDirectory string = "adr"
)

func initialize(cfg config.Configuration) {
	if err := pkg.InitializeProject(cfg); err != nil {
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

	cfg, err := pkg.LoadConfiguration()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	initialize(cfg)

	adr := Adr{
		title: "Use ADRs",
		status: []AdrStatus{
			AdrStatus{kind: Drafted, date: "2022-11-05"},
			AdrStatus{kind: Proposed, date: "2022-11-05"},
		},
		context:      "Keeping track of *why* something was done gets error-prone over time.",
		decision:     "We're gonna start using ADRs to document major decisions",
		consequences: "Start doctools utility development to facilitate this",
	}
	dbg.Debug("\n\n%v\n\n", adr)
}
