package main

import (
	"strings"
)

type wpPlugin struct {
	name        string
	description string
	plugin      string
}

func projectPhpPlugin(p projectInfo) wpPlugin {
	var data wpPlugin

	for _, fname := range p.listFiles("*.php") {
		raw, err := p.getFile(fname)
		if err != nil {
			continue
		} else if !strings.Contains(string(raw), "Plugin Name:") {
			continue
		}
		data.plugin = fname

		lines := strings.Split(string(raw), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Plugin Name:") {
				parts := strings.Split(line, "Plugin Name:")
				data.name = strings.TrimSpace(parts[len(parts)-1])
			}
			if strings.Contains(line, "Description:") {
				parts := strings.Split(line, "Description:")
				data.description = strings.TrimSpace(parts[len(parts)-1])
				if data.name != "" {
					break
				}
			}
		}
		if data.plugin != "" {
			break
		}
	}

	return data
}
