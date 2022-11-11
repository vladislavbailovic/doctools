package main

import "testing"

func Test_getWIPChanges(t *testing.T) {
	res := getWIPChanges()
	if len(res) == 0 {
		t.Fatal("expected some commits here, got none")
	}
}

func Test_getChangesets(t *testing.T) {
	res := getChangesets()
	t.Fatalf("changesets: %#v", res)
}
