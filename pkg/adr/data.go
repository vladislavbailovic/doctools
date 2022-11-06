package adr

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

//go:embed resources/template.md
var templateSource string
var tpl = template.Must(
	template.New("ADR").Parse(templateSource),
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
	buffer := new(strings.Builder)
	tpl.Execute(buffer, x)
	return buffer.String()
}

func parseData(raw string) (Data, error) {
	var data Data
	lines := strings.Split(raw, "\n")
	if len(lines) < 10 {
		return data, ParseError("input too short")
	}

	// title & number
	parts := strings.SplitN(lines[0], ":", 2)
	if len(parts) != 2 {
		return data, ParseError("unrecognized title format: %s", lines[0])
	}
	data.Title = strings.TrimSpace(parts[1])
	// number
	number, err := strconv.Atoi(strings.TrimSpace(parts[0][3:]))
	if err != nil {
		return data, ParseError("unrecognized number format (%s): %w", parts[0], err)
	}
	data.Number = uint(number)

	lineIdx, err := next(lines, 1, "Status")
	if err != nil {
		return data, ParseError("no status: %w", err)
	}

	// status
	for _, rawStatus := range strings.Split(lines[lineIdx], ",") {
		status, err := parseStatus(strings.TrimSpace(rawStatus))
		if err != nil {
			return data, ParseError("invalid status: %w", err)
		}
		data.Status = append(data.Status, status)
	}

	// context
	lineIdx, err = next(lines, lineIdx, "Context")
	if err != nil {
		return data, ParseError("invalid context: %w", err)
	}
	nextIdx, err := next(lines, lineIdx, "Decision")
	if err != nil {
		return data, ParseError("invalid decision: %w", err)
	}
	data.Context = strings.TrimSpace(strings.Join(lines[lineIdx:nextIdx-3], "\n"))
	lineIdx = nextIdx

	// decision
	nextIdx, err = next(lines, lineIdx, "Consequences")
	if err != nil {
		return data, ParseError("invalid consequences: %w", err)
	}
	data.Decision = strings.TrimSpace(strings.Join(lines[lineIdx:nextIdx-3], "\n"))
	lineIdx = nextIdx

	// consequences
	data.Consequences = strings.TrimSpace(strings.Join(lines[lineIdx:], "\n"))

	return data, nil
}

func next(lines []string, from int, pattern string) (int, error) {
	lineIdx := 1
	for lines[lineIdx] != pattern {
		lineIdx++
		if lineIdx > len(lines) {
			break
		}
	}
	if len(lines) < lineIdx+3 {
		return lineIdx, ParseError("expected %s at %d", pattern, lineIdx)
	}
	return lineIdx + 3, nil
}

func ParseError(msg string, rest ...interface{}) error {
	return fmt.Errorf(
		fmt.Sprintf("parse error: %s", msg),
		rest...)
}
