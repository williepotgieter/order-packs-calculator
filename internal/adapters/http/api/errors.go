package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// httpError Handles errors that might occur in handler functions in a standard way.
func httpError(c *gin.Context, logger *zap.Logger, code int, message string, err error) {
	var level zapcore.Level

	// Assign log level based on the HTTP response code
	switch {
	case code > 399 && code < 500:
		level = zap.WarnLevel
	case code >= 500:
		level = zap.ErrorLevel

	default:
		level = zap.DebugLevel
	}

	logger.Log(level, message, zap.Error(err))

	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
