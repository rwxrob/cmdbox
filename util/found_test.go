package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFoundPathRequired(t *testing.T) {
	if Found("") == true {
		t.Fatal("Found(\"\") should return false")
	}
}

func TestFound(t *testing.T) {
	dir, err := os.MkdirTemp("", "cmdbox")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	filenames := []string{"file.go", "config.json"}
	for _, filename := range filenames {
		path := filepath.Join(dir, filename)
		_, err := os.Create(path)
		if err != nil {
			t.Fatal(err)
		}

		if Found(path) == false {
			t.Fatalf("Expected Found(%s) to return true\n", filename)
		}
	}
}

func TestNotFound(t *testing.T) {
	filenames := []string{"hhtp.go", "so.go", "shh.go"}
	for _, filename := range filenames {
		if Found(filename) == true {
			t.Fatalf("Expected Found(%s) to return false", filename)
		}
	}
}
