package main

import "testing"

func Test_FactoryError(t *testing.T) {
	nfo := projectInfo{}

	if nfo.path != "" {
		t.Fatalf("invalid project should not have path: %v", nfo.path)
	}
	if nfo.getPath("whatever") != "" {
		t.Fatalf("invalid project should not resolve paths: %v", nfo.getPath("whatever"))
	}
}

func Test_HasFile(t *testing.T) {
	nfo := newProjectInfo("../..")
	if !nfo.hasFile("go.mod") {
		t.Fatalf("this project should have go.mod file: %v", nfo)
	}
}

func Test_HasDir(t *testing.T) {
	nfo := newProjectInfo("../..")
	if !nfo.hasDir("testdata") {
		t.Fatalf("this project should have testdata dir: %v", nfo)
	}
	if nfo.hasFile("testdata") {
		t.Fatalf("testdata should be a dir: %v", nfo)
	}
}

func Test_GetFile_ReturnsBytes(t *testing.T) {
	nfo := newProjectInfo("../..")

	if mod, err := nfo.getFile("go.mod"); err != nil {
		t.Fatalf("this project should have go.mod file: %v", nfo)
	} else if len(mod) == 0 {
		t.Fatalf("this go.mod file should not be empty: %v", mod)
	}
}

func Test_GetPaths_HasPaths(t *testing.T) {
	nfo := newProjectInfo("../../testdata/wp-plugin")
	list := nfo.listFiles("*.xml")
	if len(list) != 2 {
		t.Fatalf("expected two XML files, got: %v", list)
	}

	if !nfo.hasFiles("*.php") {
		t.Fatal("expected to recognize PHP file presence")
	}
}

func Test_ListFiles_RegularGlob(t *testing.T) {
	nfo := newProjectInfo("../../testdata/wp-plugin")
	list := nfo.listFiles("*.xml")
	if len(list) != 2 {
		t.Fatalf("expected two XML files, got: %v", list)
	}
}

func Test_ListFiles_DoubleStarGlob(t *testing.T) {
	nfo := newProjectInfo("../../testdata")
	list := nfo.listFiles("*/Dockerfile")
	if len(list) != 0 {
		t.Fatalf("expected zero top-level Dockerfiles without double star, got: %v", list)
	}

	list = nfo.listFiles("**/Dockerfile")
	if len(list) != 3 {
		t.Fatalf("expected three Dockerfiles with double star, got: %v", list)
	}
}
