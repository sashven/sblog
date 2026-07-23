package content

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadDirSortsNewestFirstAndIndexesBySlug(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "older-post-01-02-2025.md", "---\ntitle: Older\nsummary: Old summary\n---\n# Older\n")
	writePost(t, dir, "newer-post-07-13-2026.md", "---\ntitle: Newer\nsummary: New summary\n---\n# Newer\n")

	posts, err := LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir() error = %v", err)
	}
	if len(posts) != 2 {
		t.Fatalf("len(posts) = %d, want 2", len(posts))
	}
	if posts[0].Slug != "newer-post" {
		t.Fatalf("posts[0].Slug = %q, want newer-post", posts[0].Slug)
	}

	index := NewIndex(posts)
	got, ok := index.Get("older-post")
	if !ok {
		t.Fatal("index.Get(older-post) ok = false")
	}
	if got.Title != "Older" {
		t.Fatalf("got.Title = %q, want Older", got.Title)
	}
}

func TestLoadDirRejectsDuplicatePreviewNames(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "weekly-update-07-01-2026.md", "---\ntitle: First\n---\n# First\n")
	writePost(t, dir, "weekly-update-07-08-2026.md", "---\ntitle: Second\n---\n# Second\n")

	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("LoadDir() error = nil, want duplicate preview-name error")
	}
	if !strings.Contains(err.Error(), `duplicate post preview name "weekly-update"`) {
		t.Fatalf("LoadDir() error = %q, want duplicate preview-name error", err.Error())
	}
}

func writePost(t *testing.T, dir string, name string, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
		t.Fatalf("write post %s: %v", name, err)
	}
}
