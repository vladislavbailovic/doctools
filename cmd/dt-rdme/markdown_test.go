package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func getTestFile(fname string) string {
	path := filepath.Join("..", "..", fname)
	cnt, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("reading %s: %v", path, err))
	}
	return string(cnt)
}

func Test_md_isHashHeading(t *testing.T) {
	if isHashHeader("wat") {
		t.Fatal("wat is not hash heading")
	}
	if !isHashHeader("# wat") {
		t.Fatal("wat is hash heading, level 1")
	}
	if !isHashHeader("## wat") {
		t.Fatal("wat is hash heading, level 2")
	}
	if !isHashHeader("### wat") {
		t.Fatal("wat is hash heading, level 3")
	}
}

func Test_md_isDashHeading(t *testing.T) {
	if isDashHeader([]string{"wat"}) {
		t.Fatal("wat is not hash heading, limit")
	}
	if isDashHeader([]string{"wat", "wat"}) {
		t.Fatal("wat is not hash heading, no dashes")
	}
	if isDashHeader([]string{"wat", "<!--ts-->"}) {
		t.Fatal("wat is not hash heading, HTML comment")
	}
	if isDashHeader([]string{"too few", "==="}) {
		t.Fatal("wat is not hash heading, too few dashes")
	}
	if !isDashHeader([]string{"heading 1", "===="}) {
		t.Fatal("wat is hash heading, level 1")
	}
	if !isDashHeader([]string{"heading 1", "====="}) {
		t.Fatal("wat is hash heading, level 1, uneven")
	}
	if !isDashHeader([]string{"heading 2", "----"}) {
		t.Fatal("wat is hash heading, level 2")
	}
	if !isDashHeader([]string{"heading 1", "-----"}) {
		t.Fatal("wat is hash heading, level 1, uneven")
	}
}

func Test_md_getHashHeaderLevel(t *testing.T) {
	suite := map[string]HeaderLevel{
		"test":       HeaderAny,
		"# test":     HeaderLevel1,
		"# t#e#s#t#": HeaderLevel1,
		"# test #":   HeaderLevel1,
		"# #test #":  HeaderLevel1,

		"## test":     HeaderLevel2,
		"### test":    HeaderLevel3,
		"#### test":   HeaderLevel4,
		"##### test":  HeaderLevel5,
		"###### test": HeaderLevel6,
	}
	for test, expected := range suite {
		actual := getHashHeaderLevel(test)
		if actual != expected {
			t.Fatalf("%s: expected %v, but got %v", test, expected, actual)
		}
	}
}

func Test_md_getDashHeaderLevel(t *testing.T) {
	suite := map[string]HeaderLevel{
		"===":  HeaderAny,
		"---":  HeaderAny,
		"====": HeaderLevel1,
		"----": HeaderLevel2,
	}
	for test, expected := range suite {
		actual := getDashHeaderLevel([]string{"wat", test})
		if actual != expected {
			t.Fatalf("%s: expected %v, but got %v", test, expected, actual)
		}
	}
}

func Test_md_getMarkdownHeaders_HashHeaders(t *testing.T) {
	lines := strings.Split(getTestFile("testdata/example-markdown/README-vimspector.md"), "\n")
	headers := getMarkdownHeaders(lines)

	if len(headers) != 96 {
		t.Fatalf("expected exactly 96 headers, got %d", len(headers))
	}

	for idx, header := range headers {
		if !isHashHeader(lines[header.pos]) {
			t.Fatalf("expected header at %d to be a hash header: %s", idx, lines[header.pos])
		}
		if isHashHeader(lines[header.contentPos]) {
			t.Fatalf("expected content at %d to not be a hash header: %s", idx, lines[header.contentPos])
		}
		if header.kind != HashHeader {
			t.Fatalf("expected header at %d to be Hash header", idx)
		}
	}
}

func Test_md_getMarkdownHeaders_DashHeaders(t *testing.T) {
	lines := strings.Split(getTestFile("testdata/adr-001.md"), "\n")
	headers := getMarkdownHeaders(lines)

	if len(headers) != 5 {
		t.Fatalf("expected exactly 5 headers, got %d", len(headers))
	}

	for idx, header := range headers {
		if !isDashHeader([]string{lines[header.pos], lines[header.pos+1]}) {
			t.Fatalf("expected header at %d to be a hash header: %s", idx, lines[header.pos])
		}
		if header.level != Heading && header.level != Subheading {
			t.Fatalf("expected level error for header at %d: %#v (%s)", idx, header, lines[header.pos])
		}
		if header.kind != DashHeader {
			t.Fatalf("expected header at %d to be dash header", idx)
		}
	}
}

func Test_md_detectMarkdownTOC(t *testing.T) {
	lines := strings.Split(getTestFile("testdata/example-markdown/generated-1.md"), "\n")
	result := detectMarkdownTOC(lines)

	if result[0] != 4 && result[1] != 10 {
		t.Fatalf("error detecting TOC in generated: %#v", result)
	}

	lines = strings.Split(getTestFile("testdata/example-markdown/README-vimspector.md"), "\n")
	result = detectMarkdownTOC(lines)

	if result[0] != 125 && result[1] != 125 {
		t.Fatalf("error detecting TOC in generated: %#v", result)
	}
}

func Test_md_markdownHeaderText(t *testing.T) {
	suite := map[string]string{
		"# test":    "test",
		"## test#":  "test#",
		"# test ##": "test ##",
	}
	for test, expected := range suite {
		actual := markdownHeaderText(test)
		if actual != expected {
			t.Fatalf("[%s]: expected [%s], got [%s]", test, expected, actual)
		}
	}
}

func Test_md_slugifyMarkdownHeader(t *testing.T) {
	suite := map[string]string{
		"C++":         "c",
		"test / test": "test--test",
	}
	for test, expected := range suite {
		actual := slugifyMarkdownHeader(test)
		if actual != expected {
			t.Fatalf("[%s]: expected [%s], got [%s]", test, expected, actual)
		}
	}
}

func Test_md_replaceMarkdownTOC(t *testing.T) {
	// lines := strings.Split(getTestFile("testdata/example-markdown/generated-1.md"), "\n")
	lines := strings.Split(getTestFile("testdata/example-markdown/README-vimspector.md"), "\n")
	result := replaceMarkdownTOC(lines)

	t.Log(strings.Join(result, "\n"))

	t.Fatal("EOTEST")

	// lines = strings.Split(getTestFile("testdata/example-markdown/README-vimspector.md"), "\n")
	// result = detectMarkdownTOC(lines)

	// if result[0] != 125 && result[1] != 125 {
	// 	t.Fatalf("error detecting TOC in generated: %#v", result)
	// }
}
