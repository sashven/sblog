package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/sashven/sblog/internal/content"
)

func TestListPosts(t *testing.T) {
	index := content.NewIndex([]content.Post{{
		Slug:        "hello-world",
		Title:       "Hello World",
		Summary:     "First post",
		Tags:        []string{"go"},
		PublishedAt: time.Date(2026, 7, 13, 0, 0, 0, 0, time.UTC),
	}})
	router := newRouter(index)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "hello-world") {
		t.Fatalf("body %q doesn't contain slug", rec.Body.String())
	}
}

func TestGetPostBySlug(t *testing.T) {
	index := content.NewIndex([]content.Post{{
		Slug:     "hello-world",
		Title:    "Hello World",
		Summary:  "First post",
		BodyHTML: "<h1>Hello World</h1>",
	}})
	router := newRouter(index)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts/hello-world", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "<h1>Hello World</h1>") {
		t.Fatalf("body %q doesn't contain rendered HTML", rec.Body.String())
	}
}

func TestGetMissingPostReturns404(t *testing.T) {
	router := newRouter(content.NewIndex(nil))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts/missing", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusNotFound, rec.Body.String())
	}
}
