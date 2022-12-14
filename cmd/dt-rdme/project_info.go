package main

import (
	"doctools/pkg/cli"
	"os"
	"path/filepath"
	"strings"
)

const doubleStarMaxGlobDepth int = 5

type projectInfo struct {
	path string
}

func (x projectInfo) hasFiles(expr string) bool {
	return len(x.listFiles(expr)) > 0
}

func (x projectInfo) listFiles(expr string) []string {
	if strings.Contains(expr, "**") {
		result := x._listFiles(strings.Replace(expr, "**", "", 1))
		levels := make([]string, 0, doubleStarMaxGlobDepth)
		for i := 0; i < doubleStarMaxGlobDepth; i++ {
			levels = append(levels, "*")
			rpl := filepath.Join(levels...)
			result = append(result, x._listFiles(strings.Replace(expr, "**", rpl, 1))...)
		}
		return result
	}
	return x._listFiles(expr)
}

func (x projectInfo) _listFiles(expr string) []string {
	var result []string
	if list, err := filepath.Glob(x.getPath(expr)); err == nil {
		for _, f := range list {
			if relpath, err := filepath.Rel(x.path, f); err == nil {
				result = append(result, relpath)
			}
		}
	}
	return result
}

func (x projectInfo) hasFile(path string) bool {
	if x.path == "" {
		return false
	}
	if nfo, err := os.Stat(x.getPath(path)); err != nil {
		return false
	} else {
		return !nfo.IsDir()
	}
}

func (x projectInfo) hasDir(path string) bool {
	if x.path == "" {
		return false
	}
	if nfo, err := os.Stat(x.getPath(path)); err != nil {
		return false
	} else {
		return nfo.IsDir()
	}
}

func (x projectInfo) getPath(path string) string {
	if x.path == "" {
		return ""
	}
	return filepath.Join(x.path, path)
}

func (x projectInfo) getFile(path string) ([]byte, error) {
	return os.ReadFile(x.getPath(path))
}

func newProjectInfo(path string) projectInfo {
	current := projectInfo{}

	if path, err := filepath.Abs(path); err != nil {
		cli.Cry("%v", err)
		return current
	} else {
		current.path = path
	}

	return current
}
