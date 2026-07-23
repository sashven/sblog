package api

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sashven/sblog/internal/content"
	"github.com/sashven/sblog/internal/web"
)

func newRouter(index content.Index) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "Sblog",
		})
	})

	api := router.Group("/api/v1")
	api.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "Sblog"})
	})
	api.GET("/posts", listPostsHandler(index))
	api.GET("/posts/:slug", getPostHandler(index))

	registerFrontend(router)
	return router
}

func registerFrontend(router *gin.Engine) {
	dist, err := fs.Sub(web.Assets, "dist")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(dist))
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/api" || strings.HasPrefix(path, "/api/") {
			notFound(c)
			return
		}
		if path == "/" || !strings.Contains(path, ".") {
			c.Request.URL.Path = "/"
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	})
}
