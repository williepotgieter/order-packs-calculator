package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestHandleCalculatePacks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := []struct {
		name         string
		payload      calculatePacksRequest
		expectStatus int
	}{
		{
			name: "happy path with correct response and code",
			payload: calculatePacksRequest{
				Items: 500000,
				Packs: []int{23, 31, 53},
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "bad request with correct code",
			payload: calculatePacksRequest{
				Packs: []int{23, 31, 53},
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			core, observedLogs := observer.New(zapcore.DebugLevel)
			logger := zap.New(core)

			bodyBytes, err := json.Marshal(tt.payload)
			if err != nil {
				t.Fatalf("unable to marshal request body JSON: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			c.Request = req

			handler := handleCalculatePacks(logger)
			handler(c)

			assert.Equal(t, tt.expectStatus, w.Code)

			logs := observedLogs.All()

			if w.Code == http.StatusOK {
				var resp calculatePacksResponse
				err := json.Unmarshal(w.Body.Bytes(), &resp)

				assert.NoError(t, err)
				assert.Equal(t, len(logs), 0)
			} else {
				assert.Equal(t, len(logs), 1)
			}
		})
	}
}
