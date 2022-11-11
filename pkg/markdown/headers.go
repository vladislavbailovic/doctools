package markdown

import (
	"strings"
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

func (x HeaderLevel) String() string {
	return strings.Repeat("#", int(x))
}

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
	switch strings.Count(strings.TrimSuffix(line, GetHeaderText(line)), "#") {
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

func GetHeaderText(ttl string) string {
	return strings.TrimSpace(strings.TrimLeft(ttl, "#"))
}

func Delistify(item string) string {
	li := strings.TrimSpace(item)
	if !strings.HasPrefix(li, "- ") && !strings.HasPrefix(li, "* ") {
		return item
	}
	return strings.TrimPrefix(strings.TrimPrefix(li, "- "), "* ")
}
