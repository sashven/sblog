package web

import "embed"

// Assets contains the built Vue frontend.
//
//go:embed all:dist
var Assets embed.FS
