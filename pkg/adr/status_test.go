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

func Test_TypeFromString_Error(t *testing.T) {
	kind, err := StatusTypeFromString("whatever")
	if err == nil {
		t.Fatalf("expected error, got %v", kind)
	}
}

func Test_TypeFromString_Draft(t *testing.T) {
	kind, err := StatusTypeFromString("draft")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Drafted {
		t.Fatalf("expected Drafted, got %v", kind)
	}

	kind, err = StatusTypeFromString("drafted")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Drafted {
		t.Fatalf("expected Drafted, got %v", kind)
	}

	kind, err = StatusTypeFromString("-d")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Drafted {
		t.Fatalf("expected Drafted, got %v", kind)
	}
}

func Test_TypeFromString_Proposal(t *testing.T) {
	kind, err := StatusTypeFromString("propose")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Proposed {
		t.Fatalf("expected Proposed, got %v", kind)
	}

	kind, err = StatusTypeFromString("proposed")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Proposed {
		t.Fatalf("expected Proposed, got %v", kind)
	}

	kind, err = StatusTypeFromString("-p")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Proposed {
		t.Fatalf("expected Drafted, got %v", kind)
	}
}

func Test_TypeFromString_Accept(t *testing.T) {
	kind, err := StatusTypeFromString("accept")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Accepted {
		t.Fatalf("expected Accepted, got %v", kind)
	}

	kind, err = StatusTypeFromString("accepted")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Accepted {
		t.Fatalf("expected Accepted, got %v", kind)
	}

	kind, err = StatusTypeFromString("-a")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Accepted {
		t.Fatalf("expected Accepted, got %v", kind)
	}
}

func Test_TypeFromString_Reject(t *testing.T) {
	kind, err := StatusTypeFromString("reject")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Rejected {
		t.Fatalf("expected Rejected, got %v", kind)
	}

	kind, err = StatusTypeFromString("rejected")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Rejected {
		t.Fatalf("expected Rejected, got %v", kind)
	}

	kind, err = StatusTypeFromString("-r")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Rejected {
		t.Fatalf("expected Rejected, got %v", kind)
	}
}

func Test_TypeFromString_Supersede(t *testing.T) {
	kind, err := StatusTypeFromString("supersede")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Superseded {
		t.Fatalf("expected Superseded, got %v", kind)
	}

	kind, err = StatusTypeFromString("superseded")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Superseded {
		t.Fatalf("expected Superseded, got %v", kind)
	}

	kind, err = StatusTypeFromString("-s")
	if err != nil {
		t.Fatal(err)
	}
	if kind != Superseded {
		t.Fatalf("expected Superseded, got %v", kind)
	}
}
