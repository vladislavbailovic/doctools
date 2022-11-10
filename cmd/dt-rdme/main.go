package main

import (
	"doctools/pkg/cli"
	"doctools/pkg/markdown"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed resources/readme.md
var templateSource string
var tpl = template.Must(
	template.New("README").Funcs(template.FuncMap{
		"slugify": markdown.SlugifyHeader,
	}).Parse(templateSource),
)

func main() {
	if !cli.HasSubcommand() {
		cli.Say("HALP!")
	} else {
		proj := newProjectInfo(".")
		switch cli.Subcommand() {
		case "new", "init":
			_, err := proj.getFile("README.md")
			if err == nil && !cli.HasFlag("-f") && !cli.HasFlag("--force") {
				cli.Cry("README.md already exists")
				cli.Say("You can forcefully overwrite it, though (-f/--force)")
				return
			}

			if err := initReadme(proj); err != nil {
				cli.Cry("%v", err)
			}
		case "update", "toc":
			path, err := proj.getFile("README.md")
			if err != nil {
				if err := initReadme(proj); err != nil {
					cli.Cry("%v", err)
					return
				}
			}
			if err := updateReadmeToc(string(path)); err != nil {
				cli.Cry("%v", err)
			}
		default:
			cli.Say("HALP!")
		}
	}
}

func initReadme(p projectInfo) error {
	nfo, err := detectProjectMeta(p)
	if err != nil {
		return err
	}

	buffer := new(strings.Builder)
	tpl.Execute(buffer, nfo)
	if err := os.WriteFile("README.md", []byte(buffer.String()), 0622); err != nil {
		return err
	}
	return nil
}

func updateReadmeToc(path string) error {
	md := markdown.NewMarkdownFromSource(path)
	updated := md.UpdateTOC()
	if err := os.WriteFile("README.md", []byte(updated.String()), 0622); err != nil {
		return err
	}
	return nil
}

func detectProjectMeta(p projectInfo) (readme, error) {
	result := readme{}
	basedir := filepath.Base(p.path)

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
	}

	if p.hasFile("go.mod") {
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

	if p.hasFiles("*.php") {
		plug := projectPhpPlugin(p)
		if plug.plugin != "" {
			result.Name = plug.name
			result.Description = plug.description

			result.addSection(newBuildSection(fmt.Sprintf("wp plugin activate %s/%s", basedir, plug.plugin)))
		}
		if p.hasFile("phpunit.xml") {
			result.addSection(newTestSection("phpunit -c phpunit.xml"))
		}
		if p.hasFile("phpcs.ruleset.xml") {
			result.addSection(newTestSection("phpcs $(find . -type f -name '*.php') --standard=./phpcs.ruleset.xml"))
		}
	}

	if p.hasFiles("**/Dockerfile") {
		for _, dockerfile := range p.listFiles("**/Dockerfile") {
			for _, sect := range newDockerSections(dockerfile) {
				result.addSection(sect)
			}
		}
	}

	if result.Name == "" {
		result.Name = basedir
	}

	return result, nil
}
