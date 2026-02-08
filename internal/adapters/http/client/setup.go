package client

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var publicFs embed.FS

//go:embed index.html
var indexHTML string

// Setup assigns HTTP handler functions to the appropriate client route
func Setup(router *gin.Engine) error {
	// Create a sub-filesystem from the embedded public FS to avoid
	// having the word "public" in the path of static files.
	subFS, err := fs.Sub(publicFs, "public")
	if err != nil {
		return err
	}

	// Path for hosting all required static files for the UI
	router.StaticFS("/assets", http.FS(subFS))

	// Render the index.html file at /
	router.GET("/", func(c *gin.Context) {
		c.Data(200,
			"text/html; charset=utf-8",
			[]byte(indexHTML),
		)
	})

	return nil
}
