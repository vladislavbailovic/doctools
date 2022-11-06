package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func createNewAdr(args []string) {
	cfg, err := config.Load()
	if err != nil {
		dbg.Error("%v", err)
		return
	}
	create(cfg, strings.Join(args, " "))
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

	editCreated(data, repo)
}

func editCreated(data adr.Data, repo adr.Repository) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return
	}
	fmt.Println(editor, repo.PathToADR(data))
	cmd := exec.Command(editor, repo.PathToADR(data))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		dbg.Debug("error starting: %v", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		dbg.Debug("error executing: %v", err)
		return
	}
}
