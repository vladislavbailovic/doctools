package pkg

import (
	"doctools/pkg/config"
	"doctools/pkg/profile"
	"doctools/pkg/project"
)

func LoadConfiguration() (config.Configuration, error) {
	if config, err := project.Load(); err == nil {
		return config, nil
	}
	if config, err := profile.Load(); err == nil {
		return config, nil
	}
	return profile.Initialize()
}

func InitializeProject(cfg config.Configuration) error {
	return project.Initialize(cfg)
}
