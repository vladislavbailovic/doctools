package adr

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_DataFromString_Title_Error(t *testing.T) {
	_, err := getParsedData("no-such-file")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "too short") {
		t.Fatalf("expected input too short error, got: %v", err)
	}
}

func Test_DataFromString_Title(t *testing.T) {
	data, err := getParsedData("adr-001.md")
	if err != nil {
		t.Fatal(err)
	}

	if data.Title != "New ADR for testing" {
		t.Fatalf("expected string from adr-001, got: %v", data.Title)
	}
}

func Test_DataFromString_Status(t *testing.T) {
	data, err := getParsedData("adr-001.md")
	if err != nil {
		t.Fatal(err)
	}

	if len(data.Status) != 1 {
		t.Fatalf("expected array from adr-001, got: %#v", data.Status)
	}

	if data.Status[0].Date != "today" {
		t.Fatalf("expected string from adr-001, got: %v", data.Status[0].Date)
	}

	if data.Status[0].Kind != Drafted {
		t.Fatalf("expected string from adr-001, got: %v", data.Status[0].Kind)
	}
}

func Test_DataFromString_MultipleStatus(t *testing.T) {
	data, err := getParsedData("adr-002.md")
	if err != nil {
		t.Fatal(err)
	}

	if len(data.Status) != 3 {
		t.Fatalf("expected array from adr-002, got: %#v", data.Status)
	}

	if data.Status[0].Date != "today" {
		t.Fatalf("expected string from adr-002, got: %v", data.Status[0].Date)
	}

	if data.Status[0].Kind != Drafted {
		t.Fatalf("expected string from adr-002, got: %v", data.Status[0].Kind)
	}

	if data.Status[1].Date != "2022-11-06" {
		t.Fatalf("expected string from adr-002, got: %v", data.Status[0].Date)
	}

	if data.Status[1].Kind != Proposed {
		t.Fatalf("expected string from adr-002, got: %v", data.Status[0].Kind)
	}

	if data.Status[2].Date != "2022-11-06" {
		t.Fatalf("expected string from adr-002, got: %v", data.Status[0].Date)
	}

	if data.Status[2].Kind != Accepted {
		t.Fatalf("expected string from adr-002, got: %v", data.Status[0].Kind)
	}
}

func Test_DataFromString_Context(t *testing.T) {
	data, err := getParsedData("adr-001.md")
	if err != nil {
		t.Fatal(err)
	}

	if data.Context != "A context line.\n\nContext can actually be multiple lines." {
		t.Fatalf("expected string from adr-001, got: %#v", data.Context)
	}
}

func Test_DataFromString_Decision(t *testing.T) {
	data, err := getParsedData("adr-001.md")
	if err != nil {
		t.Fatal(err)
	}

	if data.Decision != "Lol this again.\nSodales lacus sed a efficitur arcu ut pellentesque purus proin. A tempus arcu. Massa non accumsan. At sodales lacus sed a efficitur arcu ut pellentesque. Sed a mollis." {
		t.Fatalf("expected string from adr-001, got: %#v", data.Decision)
	}
}

func Test_DataFromString_Consequences(t *testing.T) {
	data, err := getParsedData("adr-001.md")
	if err != nil {
		t.Fatal(err)
	}

	if data.Consequences != "Whole lotta these" {
		t.Fatalf("expected string from adr-001, got: %#v", data.Consequences)
	}
}

func getParsedData(fname string) (Data, error) {
	return parseData(getTestContent(fname))
}

func getTestContent(fname string) string {
	cwd, _ := os.Getwd()
	cnt, _ := os.ReadFile(filepath.Join(cwd, "..", "..", "test-data", fname))
	return string(cnt)
}
