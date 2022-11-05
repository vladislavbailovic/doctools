package main

import (
	"fmt"
	"strings"
)

type AdrStatusType string

const (
	Drafted    AdrStatusType = "Drafted"
	Proposed   AdrStatusType = "Proposed"
	Accepted   AdrStatusType = "Accepted"
	Rejected   AdrStatusType = "Rejected"
	Superseded AdrStatusType = "Superseded"
)

type AdrStatus struct {
	kind AdrStatusType
	date string
}

func (x AdrStatus) String() string {
	kind := x.kind
	if kind == "" {
		kind = Drafted
	}

	date := x.date
	if date == "" {
		date = "Today" // TODO: fix
	}

	return fmt.Sprintf("%s(%s)", kind, date)
}

func parseStatus(s string) (AdrStatus, error) {
	var status AdrStatus

	if strings.Contains(s, string(Drafted)) {
		status.kind = Drafted
	} else if strings.Contains(s, string(Proposed)) {
		status.kind = Proposed
	} else if strings.Contains(s, string(Accepted)) {
		status.kind = Accepted
	} else if strings.Contains(s, string(Rejected)) {
		status.kind = Rejected
	} else if strings.Contains(s, string(Superseded)) {
		status.kind = Superseded
	} else {
		return status, fmt.Errorf("unknown status type: %s", s)
	}

	pos := strings.Index(s, string(status.kind)) + len(string(status.kind))
	if len(s[pos:]) > 1 {
		rest := s[pos:]
		if len(rest) > 2 && rest[0] == '(' {
			rest = rest[1:]
			for len(rest) > 1 && rest[0] != ')' {
				status.date += string(rest[0])
				rest = rest[1:]
			}

		}
	}

	return status, nil
}

type Adr struct {
	number       uint
	title        string
	context      string
	decision     string
	status       []AdrStatus
	consequences string
}

func (x Adr) String() string {
	stats := make([]string, len(x.status))
	for i, y := range x.status {
		stats[i] = y.String()
	}
	return strings.Join([]string{
		fmt.Sprintf("# ADR %03d: %s", x.number, x.title),
		fmt.Sprintf("## Status\n\n%s", strings.Join(stats, ", ")),
		fmt.Sprintf("## Context\n\n%s", x.context),
		fmt.Sprintf("## Decision\n\n%s", x.decision),
		fmt.Sprintf("## Consequences\n\n%s", x.consequences),
	}, "\n\n\n")
}
