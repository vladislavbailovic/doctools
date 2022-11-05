package project

import (
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Load() (config.Configuration, error) {
	dbg.Debug("Attempting to load project configuration")

	cfg := config.Configuration{
		DocPath: config.DocumentationPath,
	}

	configFile, err := GetFilePath()
	if err != nil {
		return cfg, config.LoadError("unable to load project config: %w", err)
	}

	buffer, err := os.ReadFile(configFile)
	if err != nil {
		return cfg, config.LoadError("unable to read %s: %w", configFile, err)
	}

	if err := json.Unmarshal(buffer, &cfg); err != nil {
		return cfg, config.LoadError("parse error: %w", err)
	}

	return cfg, nil
}

func GetFilePath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return currentDir, config.LoadError("unable to determine working directory: %w", err)
	}
	configPath := filepath.Join(currentDir, fmt.Sprintf(".%s", config.ConfigFilename))

	if _, err := os.Stat(configPath); err != nil {
		return configPath, config.LoadError("%s missing: %w", configPath, err)
	}

	return configPath, nil
}

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

func Initialize(cfg config.Configuration) error {
	if err := initializeProjectConfig(cfg); err != nil {
		return config.InitError("unable to init project config: %w", err)
	}
	if err := initializeProjectDirectory(cfg); err != nil {
		return config.InitError("unable to init project directory: %w", err)
	}

	return nil
}

func initializeProjectConfig(cfg config.Configuration) error {
	configFile, err := GetFilePath()
	if err != nil {
		buffer, err := json.Marshal(cfg)
		if err != nil {
			return config.InitError("error compiling JSON: %w")
		}
		if err := os.WriteFile(configFile, buffer, 0622); err != nil {
			return config.InitError("error writing %s: %w", configFile, err)
		}
	}
	return nil
}

func initializeProjectDirectory(cfg config.Configuration) error {
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
