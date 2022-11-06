package adr

import "testing"

func Test_ParseAdrStatus_Failure(t *testing.T) {
	s, err := parseStatus("Whatever(today)")
	if err == nil {
		t.Fatalf("expected error parsing unknown status: %s", s)
	}
}

func Test_ParseAdrStatus_Proposed(t *testing.T) {
	s, err := parseStatus("Proposed(today)")
	if err != nil {
		t.Fatalf("error parsing Proposed: %v", err)
	}

	if s.Kind != Proposed {
		t.Fatalf("couldn't parse Proposed: %v", s)
	}
}

func Test_ParseAdrStatus_Proposed_Date(t *testing.T) {
	s, err := parseStatus("Proposed(not today)")
	if err != nil {
		t.Fatalf("error parsing Proposed: %v", err)
	}

	if s.Kind != Proposed {
		t.Fatalf("couldn't parse Proposed: %v", s)
	}

	if s.Date != "not today" {
		t.Fatalf("couldn't parse Proposed date: %v", s)
	}
}
