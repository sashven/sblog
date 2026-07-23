package content

import "time"

type Post struct {
	Slug         string
	PreviewName  string
	Title        string
	Summary      string
	Tags         []string
	SourcePath   string
	PublishedAt  time.Time
	BodyMarkdown string
	BodyHTML     string
}

type PostSummary struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"published_at"`
}

type PostDetail struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"published_at"`
	BodyHTML    string    `json:"body_html"`
}

func (p Post) SummaryDTO() PostSummary {
	return PostSummary{
		Slug:        p.Slug,
		Title:       p.Title,
		Summary:     p.Summary,
		Tags:        p.Tags,
		PublishedAt: p.PublishedAt,
	}
}

func (p Post) DetailDTO() PostDetail {
	return PostDetail{
		Slug:        p.Slug,
		Title:       p.Title,
		Summary:     p.Summary,
		Tags:        p.Tags,
		PublishedAt: p.PublishedAt,
		BodyHTML:    p.BodyHTML,
	}
}
