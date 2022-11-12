package config

import "strings"

type Configuration struct {
	DocPath      string
	IgnoredPaths []string
	SlugPrefix   string
}

var _configuration Configuration

func (x Configuration) IsPathIgnored(path string) bool {
	for _, ignored := range x.IgnoredPaths {
		if strings.Contains(path, ignored) {
			return true
		}
	}
	return false
}

const (
	DocumentationPath string = "docs"
	ConfigDirectory   string = "doctools"
	ConfigFilename    string = "doctools.json"
)

func Load() (Configuration, error) {
	if config, err := loadProjectConfiguration(); err == nil {
		_configuration = config
		return config, nil
	}
	if config, err := loadGlobalConfiguration(); err == nil {
		_configuration = config
		return config, nil
	}
	config, err := initializeProfile()
	if err == nil {
		_configuration = config
	}
	return config, err
}

func Get() Configuration {
	return _configuration
}

func getDefault() Configuration {
	return Configuration{
		DocPath:      DocumentationPath,
		IgnoredPaths: []string{"testdata"},
		SlugPrefix:   "markdown-header-",
	}
}
