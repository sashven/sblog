package web

import (
	"io/fs"
	"testing"
)

func TestAssetsIncludesFrontendIndex(t *testing.T) {
	index, err := fs.ReadFile(Assets, "dist/index.html")
	if err != nil {
		t.Fatalf("Assets missing dist/index.html: %v", err)
	}
	if len(index) == 0 {
		t.Fatal("dist/index.html is empty")
	}
}
