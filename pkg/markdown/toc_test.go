package markdown

import "testing"

func Test_md_SlugifyHeader(t *testing.T) {
	suite := map[string]string{
		"C++":         "c",
		"test / test": "test--test",
	}
	for test, expected := range suite {
		actual := SlugifyHeader(test)
		if actual != expected {
			t.Fatalf("[%s]: expected [%s], got [%s]", test, expected, actual)
		}
	}
}
