package config

type Configuration struct {
	DocPath string
}

const (
	DocumentationPath string = "docs"
	ConfigDirectory   string = "doctools"
	ConfigFilename    string = "doctools.json"
)

func Load() (Configuration, error) {
	if config, err := loadProjectConfiguration(); err == nil {
		return config, nil
	}
	if config, err := loadGlobalConfiguration(); err == nil {
		return config, nil
	}
	return initializeProfile()
}
