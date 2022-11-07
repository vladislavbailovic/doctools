package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/cli"
	"doctools/pkg/config"
	"os"
	"strconv"
)

func editExisting(args []string) {
	num, err := strconv.Atoi(args[0])
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

	repo, err := adr.GetRepo(cfg)
	if err != nil {
		cli.Cry("error getting adr repo: %v", err)
		return
	}
	data, err := repo.GetByNumber(uint(num))
	if err != nil {
		cli.Cry("error getting adr data: %v", err)
		return
	}

	openForEditing(data, repo)
}

func openForEditing(data adr.Data, repo adr.Repository) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return
	}
	cmd := cli.Run(editor, repo.PathToADR(data))
	if err := cmd.Wait(); err != nil {
		cli.Cry("error executing: %v", err)
		return
	}
}
