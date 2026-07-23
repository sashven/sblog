package content

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

var postFilenamePattern = regexp.MustCompile(`^([a-z0-9]+(?:-[a-z0-9]+)*)-(\d{2})-(\d{2})-(\d{4})\.md$`)

type frontMatter struct {
	Title   string   `yaml:"title"`
	Summary string   `yaml:"summary"`
	Tags    []string `yaml:"tags"`
}

func ParsePost(path string, data []byte) (Post, error) {
	name := filepath.Base(path)
	matches := postFilenamePattern.FindStringSubmatch(name)
	if matches == nil {
		return Post{}, fmt.Errorf("post filename %q must match <preview-name>-MM-DD-YYYY.md", name)
	}

	previewName := matches[1]
	publishedAt, err := time.ParseInLocation("01-02-2006", matches[2]+"-"+matches[3]+"-"+matches[4], time.UTC)
	if err != nil {
		return Post{}, fmt.Errorf("parse date from %q: %w", name, err)
	}

	meta, markdown, err := splitFrontMatter(data)
	if err != nil {
		return Post{}, err
	}

	var html bytes.Buffer
	if err := goldmark.Convert(markdown, &html); err != nil {
		return Post{}, fmt.Errorf("render markdown: %w", err)
	}

	title := strings.TrimSpace(meta.Title)
	if title == "" {
		title = strings.ReplaceAll(previewName, "-", " ")
	}

	return Post{
		Slug:         previewName,
		PreviewName:  previewName,
		Title:        title,
		Summary:      strings.TrimSpace(meta.Summary),
		Tags:         meta.Tags,
		SourcePath:   path,
		PublishedAt:  publishedAt,
		BodyMarkdown: string(markdown),
		BodyHTML:     html.String(),
	}, nil
}

func splitFrontMatter(data []byte) (frontMatter, []byte, error) {
	if !bytes.HasPrefix(data, []byte("---\n")) {
		return frontMatter{}, data, nil
	}

	rest := data[len("---\n"):]
	end := bytes.Index(rest, []byte("\n---\n"))
	if end < 0 {
		return frontMatter{}, nil, fmt.Errorf("front matter is not closed")
	}

	var meta frontMatter
	if err := yaml.Unmarshal(rest[:end], &meta); err != nil {
		return frontMatter{}, nil, fmt.Errorf("parse front matter: %w", err)
	}

	return meta, rest[end+len("\n---\n"):], nil
}
