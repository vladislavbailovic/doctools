package adr

import (
	"fmt"
	"strings"
)

type StatusType string

const (
	Drafted    StatusType = "Drafted"
	Proposed   StatusType = "Proposed"
	Accepted   StatusType = "Accepted"
	Rejected   StatusType = "Rejected"
	Superseded StatusType = "Superseded"
)

func StatusTypeFromString(str string) (StatusType, error) {
	var status StatusType
	switch strings.ToLower(str) {
	case "draft", "drafted", "-d":
		return Drafted, nil
	case "propose", "proposed", "-p":
		return Proposed, nil
	case "accept", "accepted", "-a":
		return Accepted, nil
	case "reject", "rejected", "-r":
		return Rejected, nil
	case "supersede", "superseded", "-s":
		return Superseded, nil
	default:
		return status, fmt.Errorf("unknown status type: %s", str)
	}
}

type Status struct {
	Kind StatusType
	Date string
}

func (x Status) String() string {
	kind := x.Kind
	if kind == "" {
		kind = Drafted
	}

	date := x.Date
	if date == "" {
		date = "Today" // TODO: fix
	}

	return fmt.Sprintf("%s(%s)", kind, date)
}

func (x Data) UpdateStatus(s StatusType) Data {
	for _, ss := range x.Status {
		if ss.Kind == s {
			return x
		}
	}
	x.Status = append(x.Status, Status{Kind: s, Date: "TODO"})
	return x
}

func parseStatus(s string) (Status, error) {
	var status Status

	if strings.Contains(s, string(Drafted)) {
		status.Kind = Drafted
	} else if strings.Contains(s, string(Proposed)) {
		status.Kind = Proposed
	} else if strings.Contains(s, string(Accepted)) {
		status.Kind = Accepted
	} else if strings.Contains(s, string(Rejected)) {
		status.Kind = Rejected
	} else if strings.Contains(s, string(Superseded)) {
		status.Kind = Superseded
	} else {
		return status, fmt.Errorf("unknown status type: %s", s)
	}

	pos := strings.Index(s, string(status.Kind)) + len(string(status.Kind))
	if len(s[pos:]) > 1 {
		rest := s[pos:]
		if len(rest) > 2 && rest[0] == '(' {
			rest = rest[1:]
			for len(rest) > 1 && rest[0] != ')' {
				status.Date += string(rest[0])
				rest = rest[1:]
			}

		}
	}

	return status, nil
}
