package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sashven/sblog/internal/content"
)

func TestFrontendFallbackServesIndexForBrowserRoutes(t *testing.T) {
	router := newRouter(content.NewIndex(nil))

	req := httptest.NewRequest(http.MethodGet, "/posts/hello-world", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), `id="app"`) {
		t.Fatalf("body doesn't look like Vue index.html: %s", rec.Body.String())
	}
}

func TestFrontendFallbackKeepsAPIMissesAsJSON404(t *testing.T) {
	router := newRouter(content.NewIndex(nil))

	for _, path := range []string{"/api", "/api/v1/missing"} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusNotFound {
				t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusNotFound, rec.Body.String())
			}
			if !strings.Contains(rec.Body.String(), "not_found") {
				t.Fatalf("body doesn't contain not_found error: %s", rec.Body.String())
			}
		})
	}
}
