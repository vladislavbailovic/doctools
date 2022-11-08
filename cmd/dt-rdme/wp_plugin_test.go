package main

import "testing"

func Test_FindWpPlugin_NoPlugin(t *testing.T) {
	nfo := newProjectInfo("../../testdata")
	plug := projectPhpPlugin(nfo)

	if plug.name != "" {
		t.Fatal("expected empty name")
	}
	if plug.description != "" {
		t.Fatal("expected empty description")
	}
}

func Test_FindWpPlugin_HappyPath(t *testing.T) {
	nfo := newProjectInfo("../../testdata/wp-plugin")
	plug := projectPhpPlugin(nfo)

	if plug.name != "H5P Campus" {
		t.Fatalf("could not detect plugin name, got [%v]", plug.name)
	}
	if plug.description != "H5P Campus plugin" {
		t.Fatalf("could not detect plugin description, got [%v]", plug.description)
	}
}
