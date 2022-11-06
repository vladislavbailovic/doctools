package pkg

import (
	"doctools/pkg/config"
	"doctools/pkg/project"
)

func InitializeProject(cfg config.Configuration) error {
	if err := config.InitializeProjectConfig(cfg); err != nil {
		return err
	}
	if err := project.InitializeProjectDirectory(cfg); err != nil {
		return err
	}
	return nil
}
