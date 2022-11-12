package config

import "testing"

func Test_Configuration_IsPathIgnored(t *testing.T) {
	c := Configuration{
		IgnoredPaths: []string{"testdata"},
	}
	if !c.IsPathIgnored("/testdata/wat/Dockerfile") {
		t.Fatal("testdata should be ignored")
	}
	if c.IsPathIgnored("/wat/Dockerfile") {
		t.Fatal("wat should not be ignored")
	}
}
