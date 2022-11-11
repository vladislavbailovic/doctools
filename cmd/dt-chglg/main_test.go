package main

import "testing"

func Test_getLogLines_ReturnsArray(t *testing.T) {
	lines := getLogLines("-n", "5")
	if len(lines) != 5 {
		t.Log(lines)
		t.Fatalf("expected exactly 5 log lines, got %d", len(lines))
	}
}
