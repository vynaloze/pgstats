package pgstats

import "testing"

var versionTests = []struct {
	input  string
	output string
}{
	{"11.3", "11"},
	{"10.12", "10"},
	{"9.6.13", "9.6"},
	{"9.5.17", "9.5"},
	{"9.4.22", "9.4"},
}

func TestExtractMajorVersion(t *testing.T) {
	for _, tt := range versionTests {
		actual := extractMajorVersion(tt.input)
		if actual != tt.output {
			t.Errorf("Expected '%s'; actual '%s'", tt.output, actual)
		}
	}
}
