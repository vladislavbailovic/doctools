package main

import (
	"doctools/pkg/cli"
	"encoding/json"
)

type packageJson struct {
	Name        string
	Description string
	Version     string
	License     string
	Bin         map[string]string
	Scripts     map[string]string
}

func projectPackageJson(p projectInfo) packageJson {
	var data packageJson
	raw, err := p.getFile("package.json")
	if err != nil {
		cli.Nit("%v", err)
		return data
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		cli.Nit("%v", err)
	}
	cli.Nit("data: %#v", data)
	return data
}
