package content

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func LoadDir(dir string) ([]Post, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(entries))
	seenPreviewNames := make(map[string]string, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		post, err := ParsePost(path, data)
		if err != nil {
			return nil, err
		}
		if existingPath, ok := seenPreviewNames[post.PreviewName]; ok {
			return nil, fmt.Errorf("duplicate post preview name %q in %s and %s", post.PreviewName, existingPath, path)
		}
		seenPreviewNames[post.PreviewName] = path
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i int, j int) bool {
		return posts[i].PublishedAt.After(posts[j].PublishedAt)
	})

	return posts, nil
}
