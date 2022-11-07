package config

import (
	"doctools/pkg/output"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func InitializeProject(cfg Configuration) error {
	configFile, err := getProjectConfigFilePath()
	if err != nil {
		buffer, err := json.Marshal(cfg)
		if err != nil {
			return InitError("error compiling JSON: %w")
		}
		if err := os.WriteFile(configFile, buffer, 0622); err != nil {
			return InitError("error writing %s: %w", configFile, err)
		}
	}
	return nil
}

func loadProjectConfiguration() (Configuration, error) {
	output.Nit("Attempting to load project configuration")

	cfg := Configuration{
		DocPath: DocumentationPath,
	}

	configFile, err := getProjectConfigFilePath()
	if err != nil {
		return cfg, LoadError("unable to load project config: %w", err)
	}

	buffer, err := os.ReadFile(configFile)
	if err != nil {
		return cfg, LoadError("unable to read %s: %w", configFile, err)
	}

	if err := json.Unmarshal(buffer, &cfg); err != nil {
		return cfg, LoadError("parse error: %w", err)
	}

	return cfg, nil
}

func getProjectConfigFilePath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return currentDir, LoadError("unable to determine working directory: %w", err)
	}
	configPath := filepath.Join(currentDir, fmt.Sprintf(".%s", ConfigFilename))

	if _, err := os.Stat(configPath); err != nil {
		return configPath, LoadError("%s missing: %w", configPath, err)
	}

	return configPath, nil
}
