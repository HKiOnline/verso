package verso

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestParsePath(t *testing.T) {
	t.Parallel()

	testDataDir := "tests/data"
	testFile := "changelog_challenges.md"
	filePath := filepath.Join(testDataDir, testFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Test file not found: %s", filePath)
	}

	cl, err := ParsePath(filePath)

	if err != nil {
		t.Error(err)
	}

	foundVersionCount := len(cl.Versions)
	log.Printf("%d versions found", foundVersionCount)

	expectedVersionCount := 18
	if expectedVersionCount != foundVersionCount {
		t.Fatalf("expected to find %d versions but found %d", expectedVersionCount, foundVersionCount)
	}

	// [1.1.1-beta.1+20250419.b4fa3175-3e2f-42e2-a1d3-a47a57faecfb] - 2025-04-19

	actualSemVer := cl.Versions[1]
	expectedSemVer := Semver{Major: 1, Minor: 1, Patch: 1, PreRelease: "beta.1", Build: "20250419.b4fa3175-3e2f-42e2-a1d3-a47a57faecfb"}

	date, err := time.Parse("2006-01-02", "2025-04-19")

	if err != nil {
		t.Fatalf("could not parse test time format: %s", err)
	}

	expectedSemVer.Date = date

	if expectedSemVer.Major != actualSemVer.Major {
		t.Fatalf("expected major version %d but got %d", expectedSemVer.Major, actualSemVer.Major)
	}

	if expectedSemVer.Minor != actualSemVer.Minor {
		t.Fatalf("expected minor version %d but got %d", expectedSemVer.Minor, actualSemVer.Minor)
	}

	if expectedSemVer.Patch != actualSemVer.Patch {
		t.Fatalf("expected patch version %d but got %d", expectedSemVer.Patch, actualSemVer.Patch)
	}

	if expectedSemVer.PreRelease != actualSemVer.PreRelease {
		t.Fatalf("expected pre-release version %s but got %s", expectedSemVer.PreRelease, actualSemVer.PreRelease)
	}

	if expectedSemVer.Build != actualSemVer.Build {
		t.Fatalf("expected build info %s but got %s", expectedSemVer.Build, actualSemVer.Build)
	}

	if !actualSemVer.Date.Equal(expectedSemVer.Date) {
		t.Fatalf("expected build info %s but got %s", expectedSemVer.Date.String(), actualSemVer.Date.String())
	}
}

func TestParseBytes(t *testing.T) {
	t.Parallel()

	testDataDir := "tests/data"
	testFile := "changelog_challenges.md"
	filePath := filepath.Join(testDataDir, testFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Test file not found: %s", filePath)
	}

	bytes, err := getFile(filePath)

	if err != nil {
		t.Fatalf("failed to read test file: %s", err)
	}

	cl, err := ParseBytes(bytes)

	if err != nil {
		t.Error(err)
	}

	foundVersionCount := len(cl.Versions)
	log.Printf("%d versions found", foundVersionCount)

	expectedVersionCount := 18
	if expectedVersionCount != foundVersionCount {
		t.Fatalf("expected to find %d versions but found %d", expectedVersionCount, foundVersionCount)
	}
}
