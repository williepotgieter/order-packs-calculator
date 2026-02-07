package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func httpError(c *gin.Context, code int, message string, err error) {
	log.Printf("%s: %s\n", message, err.Error())

	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
