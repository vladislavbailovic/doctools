package main

import (
	"reflect"
	"testing"
)

func Test_getLogLines_ReturnsArray_Error(t *testing.T) {
	expected := []string{}
	actual := getLogLines("--whatever", "not really valid bro")
	if len(actual) != 0 {
		t.Log(actual)
		t.Fatalf("expected zero-length result, got %d", len(actual))
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Log(actual)
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func Test_getLogLines_ReturnsArray_HappyPath(t *testing.T) {
	lines := getLogLines("-n", "5")
	if len(lines) != 5 {
		t.Log(lines)
		t.Fatalf("expected exactly 5 log lines, got %d", len(lines))
	}
}

func Test_commitFromLogLine_HappyPath(t *testing.T) {
	cmt := commitFromLogLine("test-hash test commit title")
	if cmt.hash != "test-hash" {
		t.Fatalf("expected specific hash, got [%s] instead (%v)", cmt.hash, cmt)
	}
	if cmt.title != "test commit title" {
		t.Fatalf("expected specific title, got [%s] instead (%v)", cmt.title, cmt)
	}
}

func Test_getCommits(t *testing.T) {
	commits := getCommits("-n", "5")
	if len(commits) != 5 {
		t.Log(commits)
		t.Fatalf("expected exactly 5 commits, got %d", len(commits))
	}
}

func Test_getTags(t *testing.T) {
	tags := getTags()
	if len(tags) < 1 {
		t.Log(tags)
		t.Fatalf("Expected some tags, got %#v", tags)
	}
}
