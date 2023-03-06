package internal

import (
	"embed"
	"github.com/unrolled/render"
)

//go:embed templates/*.tmpl
var embeddedTemplates embed.FS

var (
	R = render.New(render.Options{
		Directory: "templates",
		FileSystem: &render.EmbedFileSystem{
			FS: embeddedTemplates,
		},
		Extensions: []string{".html", ".tmpl"},
	})
)