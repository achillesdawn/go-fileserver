package fsutils

import "testing"

func TestCreateDir(t *testing.T) {
	err := CreateDirs()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
}
