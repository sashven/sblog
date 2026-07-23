.PHONY: test build dev-server frontend-build

GO_CACHE_DIR := $(CURDIR)/.cache/go

export GOPATH := $(GO_CACHE_DIR)/path
export GOMODCACHE := $(GO_CACHE_DIR)/mod
export GOCACHE := $(GO_CACHE_DIR)/build

test:
	go test ./...
	cd frontend && npm run test:unit -- --run

frontend-build:
	cd frontend && npm run build

build: frontend-build
	go build ./cmd/server

dev-server:
	go run ./cmd/server

clean:
	go clean
	$(RM) -r "$(GO_CACHE_DIR)"
