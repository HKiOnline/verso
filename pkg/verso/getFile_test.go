package verso

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFile_SuccessfulRead(t *testing.T) {

	testDataDir := "../../tests/data"
	testFile := "keep_a_changelog.md"
	filePath := filepath.Join(testDataDir, testFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Test file not found: %s", filePath)
	}

	fileBytes, err := getFile(filePath)
	if err != nil {
		t.Fatalf("getFile returned an error: %v", err)
	}

	if len(fileBytes) == 0 {
		t.Fatal("getFile returned an empty byte slice")
	}

}

func TestGetFile_FileNotFound(t *testing.T) {

	nonExistentFile := "non_existent_file.txt"

	_, err := getFile(nonExistentFile)

	if err == nil {
		t.Fatal("getFile should have returned an error for a non-existent file")
	}

}
