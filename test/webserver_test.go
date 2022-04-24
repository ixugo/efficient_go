package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockServer 模拟服务器
func mockServer() *httptest.Server {
	var f http.HandlerFunc
	f = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "welcome")
	}
	return httptest.NewServer(f)
}

// TestNewRequest 模拟发起请求
func TestNewRequest(t *testing.T) {
	server := mockServer()
	defer server.Close()

	req := httptest.NewRequest("GET", server.URL, nil)
	w := httptest.NewRecorder()

	server.Config.Handler.ServeHTTP(w, req)
	require.EqualValues(t, w.Result().StatusCode, http.StatusOK)
	require.EqualValues(t, "welcome", w.Body.String())
}
