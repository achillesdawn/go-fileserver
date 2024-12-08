package storage

import (
	"path/filepath"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := CreateDirs()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
}

func TestHttpDir(t *testing.T) {
	const expected string = "/home/miguel/go/file-server/storage/assets"

	path, err := filepath.Abs("./assets")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if path != expected {
		t.Fatalf("%s not equal to expected: %s", path, expected)
	}

}
