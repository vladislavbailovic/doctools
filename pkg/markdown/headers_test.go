package markdown

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

func Test_md_markdownHeaderText(t *testing.T) {
	suite := map[string]string{
		"# test":    "test",
		"## test#":  "test#",
		"# test ##": "test ##",
	}
	for test, expected := range suite {
		actual := GetHeaderText(test)
		if actual != expected {
			t.Fatalf("[%s]: expected [%s], got [%s]", test, expected, actual)
		}
	}
}

func Test_HeaderLevelToString(t *testing.T) {
	suite := map[HeaderLevel]string{
		HeaderLevel1: "#",
		HeaderLevel2: "##",
		HeaderLevel3: "###",
		HeaderLevel4: "####",
		HeaderLevel5: "#####",
		HeaderLevel6: "######",
	}
	for test, expected := range suite {
		actual := test.String()
		if actual != expected {
			t.Fatalf("expected header [%s] for level %d, got [%s]", expected, test, actual)
		}
	}
}

func Test_Listify(t *testing.T) {
	suite := []string{"", "\t", "\t\t", "\t\t\t", "\t\t\t\t", "\t\t\t\t\t"}
	for level, prefix := range suite {
		expected := prefix + "- item"
		actual := Listify("item", level)
		if expected != actual {
			t.Fatalf("expected item [%s] for level %d, got [%s]", expected, level, actual)
		}
	}
}
