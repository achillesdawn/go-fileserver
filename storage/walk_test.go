package storage

import (
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	path, err := filepath.Abs("../storage")

	t.Log("walking", path)
	if err != nil {
		panic(err)
	}
	files := Walk(path)

	for _, file := range files {
		t.Logf("%+v\n", file)
	}
}
