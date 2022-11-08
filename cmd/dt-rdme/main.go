package main

import (
	"doctools/pkg/cli"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed resources/readme.md
var templateSource string
var tpl = template.Must(
	template.New("README").Funcs(template.FuncMap{
		"slugify": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).Parse(templateSource),
)

func main() {
	if !cli.HasSubcommand() {
		cli.Say("HALP!")
	} else {
		switch cli.Subcommand() {
		case "new", "init":
			cli.Nit("Gonna create a new README file")
		case "update":
			cli.Nit("Gonna update existing readme")
			cli.Nit("This is going to be done by adding any newly detected sections and TOC")
			cli.Nit("while preserving what's already in there")
		default:
			cli.Say("HALP!")
		}
	}

	// nfo, err := getReadme("testdata")
	nfo, err := getReadme(".")
	if err != nil {
		cli.Cry("%v", err)
	}
	// cli.Say("%#v", nfo)
	buffer := new(strings.Builder)
	tpl.Execute(buffer, nfo)
	cli.Say(strings.TrimSpace(buffer.String()))
}

func getReadme(path string) (readme, error) {
	current := newProjectInfo(path)
	readme, err := detectProjectMeta(current)
	if err != nil {
		return readme, err
	}

	return readme, nil
}

func detectProjectMeta(p projectInfo) (readme, error) {
	result := readme{}

	if p.hasFile("package.json") {
		pkg := projectPackageJson(p)
		result.Name = pkg.Name
		result.Description = pkg.Description

		result.addSection(newBuildSection("npm install"))
		if _, ok := pkg.Scripts["test"]; ok {
			result.addSection(newTestSection("npm test"))
		}
		if _, ok := pkg.Scripts["build"]; ok {
			result.addSection(newBuildSection("npm run build"))
		}
		if _, ok := pkg.Scripts["start"]; ok {
			result.addSection(newRunSection("npm start"))
		}
		if pkg.Bin != nil {
			result.addSection(newBuildSection("npm link"))
			for bin, _ := range pkg.Bin {
				result.addSection(newRunSection(bin))
			}
		}
	} else if p.hasFile("go.mod") {
		mod := projectModuleFile(p)
		result.Name = mod.module

		result.addSection(newTestSection("go test ./..."))
		if len(mod.cmds) > 0 {
			result.addSection(newBuildSection(fmt.Sprintf("go build -o ./ %s/cmd/...", mod.module)))
			for _, cmd := range mod.cmds {
				result.addSection(newRunSection(fmt.Sprintf("go run %s/cmd/%s", mod.module, cmd)))
			}
		} else {
			result.addSection(newBuildSection("go build ."))
			result.addSection(newRunSection("go run ."))
		}
	}

	if result.Name == "" {
		result.Name = filepath.Base(p.path)
	}

	return result, nil
}
