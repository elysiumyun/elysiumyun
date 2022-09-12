package assets

import (
	"embed"
)

//go:embed assets/*
var Storage embed.FS
