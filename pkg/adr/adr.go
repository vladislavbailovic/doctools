package adr

import (
	"fmt"
	"strings"
)

type Data struct {
	Number       uint
	Title        string
	Context      string
	Decision     string
	Status       []Status
	Consequences string
}

func (x Data) String() string {
	stats := make([]string, len(x.Status))
	for i, y := range x.Status {
		stats[i] = y.String()
	}
	return strings.Join([]string{
		fmt.Sprintf("# ADR %03d: %s", x.Number, x.Title),
		fmt.Sprintf("## Status\n\n%s", strings.Join(stats, ", ")),
		fmt.Sprintf("## Context\n\n%s", x.Context),
		fmt.Sprintf("## Decision\n\n%s", x.Decision),
		fmt.Sprintf("## Consequences\n\n%s", x.Consequences),
	}, "\n\n\n")
}
