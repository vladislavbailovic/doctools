package adr

import (
	_ "embed"
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
