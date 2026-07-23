package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sashven/sblog/internal/content"
)

func listPostsHandler(index content.Index) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := index.List()
		summaries := make([]content.PostSummary, 0, len(posts))
		for _, post := range posts {
			summaries = append(summaries, post.SummaryDTO())
		}
		c.JSON(http.StatusOK, gin.H{"posts": summaries})
	}
}

func getPostHandler(index content.Index) gin.HandlerFunc {
	return func(c *gin.Context) {
		post, ok := index.Get(c.Param("slug"))
		if !ok {
			notFound(c)
			return
		}
		c.PureJSON(http.StatusOK, gin.H{"post": post.DetailDTO()})
	}
}
