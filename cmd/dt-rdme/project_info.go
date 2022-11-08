package main

import (
	"doctools/pkg/cli"
	"os"
	"path/filepath"
)

type projectInfo struct {
	path        string
	name        string
	description string
}

func (x projectInfo) hasFile(path string) bool {
	if "" == x.path {
		return false
	}
	if nfo, err := os.Stat(x.getPath(path)); err != nil {
		return false
	} else {
		return !nfo.IsDir()
	}
}

func (x projectInfo) hasDir(path string) bool {
	if "" == x.path {
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
