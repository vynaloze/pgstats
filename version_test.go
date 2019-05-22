package pgstats

import "testing"

var versionTests = []struct {
	input  string
	output float64
}{
	{"11.3", 11},
	{"10.12", 10},
	{"9.6.13", 9.6},
	{"9.5.17", 9.5},
	{"9.4.22", 9.4},
}

func TestExtractMajorVersion(t *testing.T) {
	for _, tt := range versionTests {
		actual, err := extractMajorVersion(tt.input)
		if err != nil {
			t.Error(err)
		}
		if actual != tt.output {
			t.Errorf("Expected '%f'; actual '%f'", tt.output, actual)
		}
	}
}
