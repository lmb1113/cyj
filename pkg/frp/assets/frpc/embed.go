package frpc

import (
	"embed"

	"cyj/pkg/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
