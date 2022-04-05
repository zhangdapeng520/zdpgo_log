package zap

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAtomicLevelServeHTTP(t *testing.T) {
	tests := []struct {
		desc          string
		method        string
		query         string
		contentType   string
		body          string
		expectedCode  int
		expectedLevel zapcore.Level
	}{
		{
			desc:          "GET",
			method:        http.MethodGet,
			expectedCode:  http.StatusOK,
			expectedLevel: zap.InfoLevel,
		},
		{
			desc:          "PUT JSON",
			method:        http.MethodPut,
			expectedCode:  http.StatusOK,
			expectedLevel: zap.WarnLevel,
			body:          `{"level":"warn"}`,
		},
		{
			desc:          "PUT URL encoded",
			method:        http.MethodPut,
			expectedCode:  http.StatusOK,
			expectedLevel: zap.WarnLevel,
			contentType:   "application/x-www-form-urlencoded",
			body:          "level=warn",
		},
		{
			desc:          "PUT query parameters",
			method:        http.MethodPut,
			query:         "?level=warn",
			expectedCode:  http.StatusOK,
			expectedLevel: zap.WarnLevel,
			contentType:   "application/x-www-form-urlencoded",
		},
		{
			desc:          "body takes precedence over query",
			method:        http.MethodPut,
			query:         "?level=info",
			expectedCode:  http.StatusOK,
			expectedLevel: zap.WarnLevel,
			contentType:   "application/x-www-form-urlencoded",
			body:          "level=warn",
		},
		{
			desc:          "JSON ignores query",
			method:        http.MethodPut,
			query:         "?level=info",
			expectedCode:  http.StatusOK,
			expectedLevel: zap.WarnLevel,
			body:          `{"level":"warn"}`,
		},
		{
			desc:         "PUT JSON unrecognized",
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			body:         `{"level":"unrecognized"}`,
		},
		{
			desc:         "PUT URL encoded unrecognized",
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			contentType:  "application/x-www-form-urlencoded",
			body:         "level=unrecognized",
		},
		{
			desc:         "PUT JSON malformed",
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			body:         `{"level":"warn`,
		},
		{
			desc:         "PUT URL encoded malformed",
			method:       http.MethodPut,
			query:        "?level=%",
			expectedCode: http.StatusBadRequest,
			contentType:  "application/x-www-form-urlencoded",
		},
		{
			desc:         "PUT Query parameters malformed",
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			contentType:  "application/x-www-form-urlencoded",
			body:         "level=%",
		},
		{
			desc:         "PUT JSON unspecified",
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			body:         `{}`,
		},
		{
			desc:         "PUT URL encoded unspecified",
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			contentType:  "application/x-www-form-urlencoded",
			body:         "",
		},
		{
			desc:         "POST JSON",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
			body:         `{"level":"warn"}`,
		},
		{
			desc:         "POST URL",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
			contentType:  "application/x-www-form-urlencoded",
			body:         "level=warn",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			lvl := zap.NewAtomicLevel()
			lvl.SetLevel(zapcore.InfoLevel)

			server := httptest.NewServer(lvl)
			defer server.Close()

			req, err := http.NewRequest(tt.method, server.URL+tt.query, strings.NewReader(tt.body))
			require.NoError(t, err, "Error constructing %s request.", req.Method)
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err, "Error making %s request.", req.Method)
			defer res.Body.Close()

			require.Equal(t, tt.expectedCode, res.StatusCode, "Unexpected status code.")
			if tt.expectedCode != http.StatusOK {
				// Don't need to test exact error message, but one should be present.
				var pld struct {
					Error string `json:"error"`
				}
				require.NoError(t, json.NewDecoder(res.Body).Decode(&pld), "Decoding response body")
				assert.NotEmpty(t, pld.Error, "Expected an error message")
				return
			}

			var pld struct {
				Level zapcore.Level `json:"level"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&pld), "Decoding response body")
			assert.Equal(t, tt.expectedLevel, pld.Level, "Unexpected logging level returned")
		})
	}
}
