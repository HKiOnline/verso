package verso

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {

	testDataDir := "../../tests/data"
	testFile := "changelog_challenges.md"
	filePath := filepath.Join(testDataDir, testFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Test file not found: %s", filePath)
	}

	cl, err := Parse(filePath)

	if err != nil {
		t.Error(err)
	}

	foundVersionCount := len(cl.Versions)
	log.Printf("%d versions found", foundVersionCount)

	expectedVersionCount := 17
	if expectedVersionCount != foundVersionCount {
		t.Fatalf("expected to find %d versions but found %d", expectedVersionCount, foundVersionCount)
	}

	log.Printf("latest version: %v+", cl.Versions[0])
	log.Printf("all versions:\n%v+", cl.Versions)
}
