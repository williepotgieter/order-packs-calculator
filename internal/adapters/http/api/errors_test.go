package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestHttpError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := []struct {
		name             string
		code             int
		message          string
		err              error
		expectResp       map[string]any
		expectCode       int
		expectLogLevel   zapcore.Level
		expectLogMessage string
	}{
		{
			name:    "happy path log warning level",
			code:    http.StatusNotFound,
			message: "YOLO",
			err:     errors.New("some error"),
			expectResp: gin.H{
				"code":    float64(http.StatusNotFound),
				"message": "YOLO",
			},
			expectCode:       http.StatusNotFound,
			expectLogLevel:   zapcore.WarnLevel,
			expectLogMessage: "YOLO",
		},
		{
			name:    "happy path log error level",
			code:    http.StatusInternalServerError,
			message: "YOLO",
			err:     errors.New("some error"),
			expectResp: gin.H{
				"code":    float64(http.StatusInternalServerError),
				"message": "YOLO",
			},
			expectCode:       http.StatusInternalServerError,
			expectLogLevel:   zapcore.ErrorLevel,
			expectLogMessage: "YOLO",
		},
		{
			name:    "happy path log debug level",
			code:    http.StatusMovedPermanently,
			message: "YOLO",
			err:     errors.New("some error"),
			expectResp: gin.H{
				"code":    float64(http.StatusMovedPermanently),
				"message": "YOLO",
			},
			expectCode:       http.StatusMovedPermanently,
			expectLogLevel:   zapcore.DebugLevel,
			expectLogMessage: "YOLO",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			core, observedLogs := observer.New(zapcore.DebugLevel)
			logger := zap.New(core)

			httpError(c, logger, tt.code, tt.message, tt.err)

			var resp map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to parse JSON: %v", err)
			}

			assert.EqualValues(t, tt.expectResp, resp)

			logs := observedLogs.All()

			assert.Equal(t, len(logs), 1)
			assert.Equal(t, tt.expectCode, w.Code)

			entry := logs[0]

			assert.Equal(t, tt.expectLogLevel, entry.Level)
			assert.Equal(t, tt.expectLogMessage, entry.Message)
		})
	}
}
