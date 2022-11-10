package markdown

import (
	"fmt"
	"strings"
	"unicode"
)

type TOCItem struct {
	caption string
	slug    string
	level   int
}

func newMarkdownHeaderTOCItem(hdLevel HeaderLevel, text string) TOCItem {
	return TOCItem{
		level:   int(hdLevel) - 1,
		caption: markdownHeaderText(text),
		slug:    SlugifyHeader(text),
	}
}

func (x TOCItem) String() string {
	return fmt.Sprintf("%s- [%s](#%s)",
		strings.Repeat("\t", x.level),
		x.caption, x.slug)
}

type TOC struct {
	items []TOCItem
}

func (x TOC) headerText() string {
	return "Table of Contents"
}

func (x TOC) Header() string {
	return fmt.Sprintf("## %s", x.headerText())
}

func (x *TOC) AddItem(item TOCItem) {
	x.items = append(x.items, item)
}

func (x TOC) IsTOCHeader(header string) bool {
	return strings.ToLower(markdownHeaderText(header)) == strings.ToLower(x.headerText())
}

func Slugify(what string) string {
	clean := ""
	for _, c := range strings.ToLower(what) {
		if !unicode.IsLower(c) && !unicode.IsNumber(c) && c != ' ' && c != '-' {
			continue
		}
		clean += string(c)
	}
	return strings.ReplaceAll(clean, " ", "-")
}

func SlugifyHeader(ttl string) string {
	return Slugify(markdownHeaderText(ttl))
}
