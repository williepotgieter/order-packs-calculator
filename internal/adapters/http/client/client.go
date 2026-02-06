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

func Setup(router *gin.Engine) error {
	subFS, err := fs.Sub(publicFs, "public")
	if err != nil {
		return err
	}

	router.StaticFS("/assets", http.FS(subFS))

	router.GET("/", func(c *gin.Context) {
		c.Data(200,
			"text/html; charset=utf-8",
			[]byte(indexHTML),
		)
	})

	return nil
}
