package profile

import (
	"doctools/pkg/config"
	"doctools/pkg/dbg"
	"encoding/json"
	"os"
	"path/filepath"
)

func Load() (config.Configuration, error) {
	dbg.Debug("Attempting to load global (profile) configuration")

	cfg := config.Configuration{
		DocPath: config.DocumentationPath,
	}

	configDir, err := GetDirectoryPath()
	if err != nil {
		return cfg, config.LoadError("unable to find profile config: %w", err)
	}
	configFile := filepath.Join(configDir, config.ConfigFilename)

	buffer, err := os.ReadFile(configFile)
	if err != nil {
		return cfg, config.LoadError("unable to read %s: %w", configFile, err)
	}

	if err := json.Unmarshal(buffer, &cfg); err != nil {
		return cfg, config.LoadError("parse error: %w", err)
	}

	return cfg, nil
}

func GetDirectoryPath() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return configPath, config.InitError("can't resolve user config dir: %w", err)
	}
	configPath = filepath.Join(configPath, config.ConfigDirectory)

	if _, err := os.Stat(configPath); err != nil {
		if err := os.Mkdir(configPath, 0722); err != nil {
			return configPath, config.InitError("config dir creation (%s): %w", configPath, err)
		}
	}

	if nfo, err := os.Stat(configPath); err != nil {
		return configPath, config.InitError("config path %s missing: %w", configPath, err)
	} else if !nfo.IsDir() {
		return configPath, config.InitError("%s is not a directory", configPath)
	}

	return configPath, nil
}

func Initialize() (config.Configuration, error) {
	dbg.Debug("Initializing global (profile) configuration")

	cfg := config.Configuration{
		DocPath: config.DocumentationPath,
	}

	base, err := GetDirectoryPath()
	if err != nil {
		return cfg, err
	}

	buffer, err := json.Marshal(cfg)
	if err != nil {
		return cfg, config.InitError("error compiling JSON: %w")
	}

	targetFile := filepath.Join(base, config.ConfigFilename)
	if err := os.WriteFile(targetFile, buffer, 0622); err != nil {
		return cfg, config.InitError("error writing %s: %w", targetFile, err)
	}

	return cfg, nil
}
