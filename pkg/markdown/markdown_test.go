package markdown

import (
	"reflect"
	"testing"
)

func Test_md_FindHeader_Heading(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/README-vimspector.md"))

	pos := md.FindHeader(Heading)
	if pos != 0 {
		t.Log(md.lines[pos])
		t.Fatalf("expected first heading at pos 0, got %d", pos)
	}
}

func Test_md_FindHeader_Subheading(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/README-vimspector.md"))

	pos := md.FindHeader(Subheading)
	if pos != 125 {
		t.Log(md.lines[pos])
		t.Fatalf("expected first subheading at pos 125, got %d", pos)
	}
}

func Test_md_FindHeaderAfter(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/README-vimspector.md"))

	positions := []int{
		0, 106, 185, 571, 593,
		684, 805, 1369, 2183, 2472,
	}
	for idx, expected := range positions {
		actual := md.FindHeaderAfter(expected-1, Heading)
		if actual != expected {
			t.Log(md.lines[actual])
			t.Fatalf("%03d: expected first heading at pos %d, got %d", idx, actual, expected)
		}
	}

	if md.FindHeaderAfter(2472, Heading) != -1 {
		t.Fatalf("expected last Heading to be last")
	}
}

func Test_md_ExtractTOC_NoToc(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/README-vimspector.md"))
	toc := md.ExtractTOC()

	if len(toc.items) != 96 {
		t.Log(toc)
		t.Fatalf("expected 96 TOC items, got %d", len(toc.items))
	}
}

func Test_md_ExtractTOC_WithToc(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/generated-1.md"))
	toc := md.ExtractTOC()

	if len(toc.items) != 5 {
		t.Log(toc)
		t.Fatalf("expected 5 TOC items, got %d", len(toc.items))
	}
}

func Test_md_HasToc_NoToc(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/README-vimspector.md"))
	if md.HasTOC() {
		t.Fatal("expected no TOC")
	}
}

func Test_md_HasToc_WithToc(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/generated-1.md"))
	if !md.HasTOC() {
		t.Fatal("expected TOC")
	}
}

func Test_md_ReplaceTOCWith_MadeUpToc(t *testing.T) {
	toc := TOC{}
	toc.AddItem(TOCItem{level: 1, caption: "test-1", slug: "wat-wat"})
	toc.AddItem(TOCItem{level: 1, caption: "test-2", slug: "wat-wat"})
	toc.AddItem(TOCItem{level: 2, caption: "sub-test-1", slug: "wat-wat"})

	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/generated-1.md"))
	updated := md.ReplaceTOCWith(toc)

	if !updated.HasTOC() {
		t.Log(updated)
		t.Fatal("TOC replacement resulted with no TOC")
	}

	newToc := updated.ExtractTOC()
	if reflect.DeepEqual(newToc, toc) {
		t.Log(updated)
		t.Fatalf("TOC replacement failure!\nexpect: %#v\nactual: %#v", toc, newToc)
	}
}

func Test_md_UpdateTOC_NoToc(t *testing.T) {
	md := NewMarkdownFromSource(getTestFile("testdata/example-markdown/README-vimspector.md"))
	updated := md.UpdateTOC()

	if !updated.HasTOC() {
		t.Log(updated)
		t.Fatal("TOC replacement resulted with no TOC")
	}

	// 	t.Log("\n--------------\n" + strings.Join(updated.lines, "\n"))
	// 	t.Fatal("dsfdsf")
}
