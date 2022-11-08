package main

import "testing"

func Test_LoadsValidPackageJson(t *testing.T) {
	p := newProjectInfo("../../testdata")
	pkg := projectPackageJson(p)
	if pkg.Name == "" {
		t.Fatalf("project %#v package.json should have a name", p)
	}
}

func Test_DoesNotLoadMissing(t *testing.T) {
	p := newProjectInfo(".")
	pkg := projectPackageJson(p)
	if pkg.Name != "" {
		t.Fatalf("there should be no package.json for project %#v", p)
	}
}
