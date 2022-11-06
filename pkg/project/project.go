package project

import (
	"os"
	"path/filepath"

	"doctools/pkg/config"
	"doctools/pkg/dbg"
)

func GetDocsDirectory(cfg config.Configuration) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return currentDir, dbg.PathError("unable to determine working directory: %w", err)
	}
	docPath := filepath.Join(currentDir, cfg.DocPath)

	if nfo, err := os.Stat(docPath); err != nil {
		return docPath, dbg.PathError("doc path %s missing: %w", docPath, err)
	} else if !nfo.IsDir() {
		return docPath, dbg.PathError("%s is not a directory", docPath)
	}

	return docPath, nil
}

func InitializeProjectDirectory(cfg config.Configuration) error {
	docPath, err := GetDocsDirectory(cfg)
	if err != nil {
		if _, err := os.Stat(docPath); err != nil {
			if err := os.Mkdir(docPath, 0722); err != nil {
				return config.InitError("doc dir creation (%s): %w", docPath, err)
			}
		} else {
			return config.InitError("doc dir resolution (%s): %w", docPath, err)
		}
	}

	if nfo, err := os.Stat(docPath); err != nil {
		return config.InitError("doc path %s missing: %w", docPath, err)
	} else if !nfo.IsDir() {
		return config.InitError("%s is not a directory", docPath)
	}

	return nil
}
