package markdown

import (
	"fmt"
	"strings"
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
		slug:    slugifyMarkdownHeader(text),
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

func replaceMarkdownTOC(lines []string) []string {
	headers := getMarkdownHeaders(lines)

	headingContent := 0
	afterHeading := 0
	tocContent := 0
	afterToc := 0
	for _, h := range headers {
		if h.level == Heading && headingContent == 0 {
			headingContent = h.contentPos
			continue
		}
		if headingContent > 0 && afterHeading == 0 {
			afterHeading = h.pos
		}
		if headingContent > 0 && tocContent > 0 {
			afterToc = h.pos
			break
		}

		if strings.Contains(strings.ToLower(lines[h.pos]), "table of content") {
			tocContent = h.pos
		}
	}

	start := 0
	end := 0
	if tocContent > 0 && afterToc > 0 {
		start = tocContent
		end = afterToc
	} else {
		start = afterHeading
		end = afterHeading
	}

	result := []string{}
	result = append(result, lines[0:start]...)
	result = append(result, "## Table of Contents")
	result = append(result, "")

	for _, h := range headers {
		if h.contentPos == headingContent {
			continue
		}
		if strings.Contains(strings.ToLower(lines[h.pos]), "table of content") {
			continue
		}

		item := newMarkdownHeaderTOCItem(h.level, lines[h.pos])
		result = append(result, item.String())
	}

	result = append(result, "")
	result = append(result, "")

	result = append(result, lines[end:]...)
	return result
}
