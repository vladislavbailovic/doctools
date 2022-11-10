package main

import (
	"fmt"
	"strings"
	"unicode"
)

type HeaderType uint8

const (
	HashHeader HeaderType = iota
	DashHeader HeaderType = iota
)

type HeaderLevel int

const (
	HeaderAny    HeaderLevel = iota
	HeaderLevel1 HeaderLevel = iota
	HeaderLevel2 HeaderLevel = iota
	HeaderLevel3 HeaderLevel = iota
	HeaderLevel4 HeaderLevel = iota
	HeaderLevel5 HeaderLevel = iota
	HeaderLevel6 HeaderLevel = iota

	Heading    = HeaderLevel1
	Subheading = HeaderLevel2
)

type headerMarker struct {
	level      HeaderLevel
	kind       HeaderType
	pos        int
	contentPos int
}

func isHashHeader(line string) bool {
	return strings.HasPrefix(line, "#")
}

func getHashHeaderLevel(line string) HeaderLevel {
	switch strings.Count(strings.TrimSuffix(line, markdownHeaderText(line)), "#") {
	case 1:
		return HeaderLevel1
	case 2:
		return HeaderLevel2
	case 3:
		return HeaderLevel3
	case 4:
		return HeaderLevel4
	case 5:
		return HeaderLevel5
	case 6:
		return HeaderLevel6
	default:
		return HeaderAny
	}
}

func isDashHeader(lines []string) bool {
	if len(lines) != 2 {
		return false
	}

	if strings.Count(lines[1], "==") > 1 {
		return len(strings.TrimSpace(strings.Replace(lines[1], "=", "", -1))) == 0
	}
	if strings.Count(lines[1], "--") > 1 {
		return len(strings.TrimSpace(strings.Replace(lines[1], "-", "", -1))) == 0
	}

	return false
}

func getDashHeaderLevel(lines []string) HeaderLevel {
	if !isDashHeader(lines) {
		return HeaderAny
	}
	if strings.Contains(lines[1], "--") {
		return Subheading
	}
	return Heading
}

func getMarkdownHeaders(lines []string) []headerMarker {
	headers := []headerMarker{}
	for idx, line := range lines {
		if idx+2 < len(lines) && isDashHeader([]string{line, lines[idx+1]}) {
			header := headerMarker{
				level:      getDashHeaderLevel([]string{line, lines[idx+1]}),
				kind:       DashHeader,
				pos:        idx,
				contentPos: idx + 2,
			}
			headers = append(headers, header)
		} else if isHashHeader(line) {
			header := headerMarker{
				level:      getHashHeaderLevel(line),
				kind:       HashHeader,
				pos:        idx,
				contentPos: idx + 1,
			}
			headers = append(headers, header)
		}
	}
	return headers
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

func markdownHeaderText(ttl string) string {
	return strings.TrimSpace(strings.TrimLeft(ttl, "#"))
}

func slugifyMarkdownHeader(ttl string) string {
	clean := ""
	for _, c := range strings.ToLower(markdownHeaderText(ttl)) {
		if !unicode.IsLower(c) && !unicode.IsNumber(c) && c != ' ' && c != '-' {
			continue
		}
		clean += string(c)
	}
	return strings.ReplaceAll(clean, " ", "-")
}
