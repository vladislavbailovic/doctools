package markdown

import (
	"strings"
)

type Markdown struct {
	lines   []string
	headers []headerMarker
}

func NewMarkdownFromLines(lines []string) Markdown {
	return Markdown{
		lines:   lines,
		headers: getMarkdownHeaders(lines),
	}
}

func NewMarkdownFromSource(src string) Markdown {
	lines := strings.Split(src, "\n")
	return NewMarkdownFromLines(lines)
}

func (x Markdown) String() string {
	return strings.TrimSpace(strings.Join(x.lines, "\n"))
}

func (x Markdown) FindHeader(lvl HeaderLevel) int {
	return x.FindHeaderAfter(-1, lvl)
}

func (x Markdown) FindHeaderAfter(line int, lvl HeaderLevel) int {
	m := x.findHeaderAfter(line, lvl)
	if m == (headerMarker{}) {
		return -1
	}
	return m.pos
}

func (x Markdown) findHeaderAfter(line int, lvl HeaderLevel) headerMarker {
	for _, h := range x.headers {
		if h.pos <= line {
			continue
		}
		if lvl == HeaderAny || h.level == lvl {
			return h
		}
	}

	return headerMarker{}
}

func (x Markdown) HasTOC() bool {
	toc := TOC{}
	for _, h := range x.headers {
		if toc.IsTOCHeader(x.lines[h.pos]) {
			return true
		}
	}
	return false
}

func (x Markdown) ExtractTOC() TOC {
	toc := TOC{}
	first := x.findHeaderAfter(-1, Heading)

	for _, h := range x.headers {
		if h.pos == first.pos {
			continue // Omit top-level heading
		}
		if toc.IsTOCHeader(x.lines[h.pos]) {
			continue
		}
		toc.AddItem(newMarkdownHeaderTOCItem(h.level, x.lines[h.pos]))
	}

	return toc
}

func (x Markdown) UpdateTOC() Markdown {
	toc := x.ExtractTOC()
	return x.ReplaceTOCWith(toc)
}

func (x Markdown) ReplaceTOCWith(toc TOC) Markdown {
	first := x.findHeaderAfter(-1, HeaderAny)
	start := first.contentPos
	end := start
	if start > 0 {
		pos := start
		for pos > 0 {
			pos = x.FindHeaderAfter(pos, HeaderAny)
			if pos > 0 && toc.IsTOCHeader(x.lines[pos]) {
				start = pos
				break
			}
		}
		if start != first.contentPos {
			end = x.FindHeaderAfter(start, HeaderAny)
		}
	}

	result := []string{}
	result = append(result, x.lines[0:start]...)
	if end == first.contentPos {
		result = append(result, "")
	}
	result = append(result, toc.Header())
	result = append(result, "")
	for _, item := range toc.items {
		result = append(result, item.String())
	}
	result = append(result, "")
	result = append(result, "")
	result = append(result, x.lines[end:]...)

	return NewMarkdownFromLines(result)
}
