package content

type Index struct {
	posts  []Post
	bySlug map[string]Post
}

func NewIndex(posts []Post) Index {
	copied := append([]Post(nil), posts...)
	bySlug := make(map[string]Post, len(copied))
	for _, post := range copied {
		bySlug[post.Slug] = post
	}
	return Index{
		posts:  copied,
		bySlug: bySlug,
	}
}

func (i Index) List() []Post {
	return append([]Post(nil), i.posts...)
}

func (i Index) Get(slug string) (Post, bool) {
	post, ok := i.bySlug[slug]
	return post, ok
}
