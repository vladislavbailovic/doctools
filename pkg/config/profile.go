package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"doctools/pkg/dbg"
)

func loadGlobalConfiguration() (Configuration, error) {
	dbg.Debug("Attempting to load global (profile) configuration")

	cfg := Configuration{
		DocPath: DocumentationPath,
	}

	configDir, err := getProfileDirectory()
	if err != nil {
		return cfg, LoadError("unable to find profile config: %w", err)
	}
	configFile := filepath.Join(configDir, ConfigFilename)

	buffer, err := os.ReadFile(configFile)
	if err != nil {
		return cfg, LoadError("unable to read %s: %w", configFile, err)
	}

	if err := json.Unmarshal(buffer, &cfg); err != nil {
		return cfg, LoadError("parse error: %w", err)
	}

	return cfg, nil
}

func initializeProfile() (Configuration, error) {
	dbg.Debug("Initializing global (profile) configuration")

	cfg := Configuration{
		DocPath: DocumentationPath,
	}

	base, err := getProfileDirectory()
	if err != nil {
		return cfg, err
	}

	buffer, err := json.Marshal(cfg)
	if err != nil {
		return cfg, InitError("error compiling JSON: %w")
	}

	targetFile := filepath.Join(base, ConfigFilename)
	if err := os.WriteFile(targetFile, buffer, 0622); err != nil {
		return cfg, InitError("error writing %s: %w", targetFile, err)
	}

	return cfg, nil
}

func getProfileDirectory() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return configPath, InitError("can't resolve user config dir: %w", err)
	}
	configPath = filepath.Join(configPath, ConfigDirectory)

	if _, err := os.Stat(configPath); err != nil {
		if err := os.Mkdir(configPath, 0722); err != nil {
			return configPath, InitError("config dir creation (%s): %w", configPath, err)
		}
	}

	if nfo, err := os.Stat(configPath); err != nil {
		return configPath, InitError("config path %s missing: %w", configPath, err)
	} else if !nfo.IsDir() {
		return configPath, InitError("%s is not a directory", configPath)
	}

	return configPath, nil
}
