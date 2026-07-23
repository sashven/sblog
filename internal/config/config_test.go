package config

import (
	"reflect"
	"testing"
)

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("CONTENT_DIR", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Env != "development" {
		t.Fatalf("Env = %q, want development", cfg.Env)
	}
	if cfg.HTTPAddr != ":8080" {
		t.Fatalf("HTTPAddr = %q, want :8080", cfg.HTTPAddr)
	}
	if cfg.ContentDir != "content/posts" {
		t.Fatalf("ContentDir = %q, want content/posts", cfg.ContentDir)
	}
}

func TestLoadReadsContentDir(t *testing.T) {
	t.Setenv("CONTENT_DIR", "testdata/posts")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.ContentDir != "testdata/posts" {
		t.Fatalf("ContentDir = %q, want testdata/posts", cfg.ContentDir)
	}
}

func TestConfigDoesNotExposeDatabaseURL(t *testing.T) {
	cfgType := reflect.TypeOf(Config{})
	if _, ok := cfgType.FieldByName("DatabaseURL"); ok {
		t.Fatal("Config exposes DatabaseURL, but the source-first MVP must not require database configuration")
	}
}
