package verso_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hkionline/verso"
)

func TestSemverString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		semver   verso.Semver
		expected string
	}{
		{
			name: "all fields",
			semver: verso.Semver{
				Major:      1,
				Minor:      1,
				Patch:      2,
				PreRelease: "beta",
				Build:      "002",
				Date:       time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: "[1.1.2-beta+002] 2023-10-01",
		},
		{
			name: "default Major.Minor.Patch",
			semver: verso.Semver{
				Major: 0,
				Minor: 0,
				Patch: 0,
			},
			expected: "",
		},
		{
			name: "empty PreRelease field",
			semver: verso.Semver{
				Major: 1,
				Minor: 1,
				Patch: 2,
				Build: "002",
				Date:  time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: "[1.1.2+002] 2023-10-01",
		},
		{
			name: "empty Build field",
			semver: verso.Semver{
				Major:      1,
				Minor:      1,
				Patch:      2,
				PreRelease: "beta",
				Date:       time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: "[1.1.2-beta] 2023-10-01",
		},
		{
			name: "default Date field",
			semver: verso.Semver{
				Major:      1,
				Minor:      1,
				Patch:      2,
				PreRelease: "beta",
				Build:      "002",
				Date:       time.Time{},
			},
			expected: "[1.1.2-beta+002]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testStringer(test.semver, test.expected, t)
		})
	}
}

func testStringer(s fmt.Stringer, expected string, t *testing.T) {
	actual := s.String()

	if actual != expected {
		t.Fatalf("expected \"%s\" but got \"%s\"\n", expected, actual)
	}
}
