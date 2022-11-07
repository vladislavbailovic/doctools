package main

import (
	"doctools/pkg/adr"
	"doctools/pkg/config"
	"doctools/pkg/output"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func editExisting(args []string) {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		output.Cry("%v", err)
		showHelp()
		return
	}

	cfg, err := config.Load()
	if err != nil {
		output.Cry("%v", err)
		return
	}

	repo, err := adr.GetRepo(cfg)
	if err != nil {
		output.Cry("error getting adr repo: %v", err)
		return
	}
	data, err := repo.GetByNumber(uint(num))
	if err != nil {
		output.Cry("error getting adr data: %v", err)
		return
	}

	openForEditing(data, repo)
}

func openForEditing(data adr.Data, repo adr.Repository) {
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
		output.Say("error starting: %v", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		output.Say("error executing: %v", err)
		return
	}
}
