package content

import (
	"testing"
	"time"
)

func TestParsePostReadsFilenameAndFrontMatter(t *testing.T) {
	input := []byte("---\ntitle: Hello World\nsummary: First Sblog post\ntags:\n  - go\n  - vue\n---\n# Hello\n\nBody text.\n")

	post, err := ParsePost("content/posts/hello-world-07-13-2026.md", input)
	if err != nil {
		t.Fatalf("ParsePost() error = %v", err)
	}
	if post.Slug != "hello-world" {
		t.Fatalf("Slug = %q, want hello-world", post.Slug)
	}
	if post.PreviewName != "hello-world" {
		t.Fatalf("PreviewName = %q, want hello-world", post.PreviewName)
	}
	wantDate := time.Date(2026, 7, 13, 0, 0, 0, 0, time.UTC)
	if !post.PublishedAt.Equal(wantDate) {
		t.Fatalf("PublishedAt = %s, want %s", post.PublishedAt, wantDate)
	}
	if post.Title != "Hello World" {
		t.Fatalf("Title = %q, want Hello World", post.Title)
	}
	if post.Summary != "First Sblog post" {
		t.Fatalf("Summary = %q, want First Sblog post", post.Summary)
	}
	if len(post.Tags) != 2 || post.Tags[0] != "go" || post.Tags[1] != "vue" {
		t.Fatalf("Tags = %#v, want [go vue]", post.Tags)
	}
	if post.BodyHTML == "" {
		t.Fatal("BodyHTML is empty")
	}
}

func TestParsePostRejectsInvalidFilename(t *testing.T) {
	_, err := ParsePost("content/posts/hello-world.md", []byte("# Hello"))
	if err == nil {
		t.Fatal("ParsePost() error is nil")
	}
}
