package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		url        string
		statusCode int
	}{
		{"http://www.baidu.com", http.StatusOK},
		{"https://ww.baadu.com", http.StatusNotFound},
	}

	for i, v := range tests {

		resp, err := http.Get(v.url)
		require.NoError(t, err)
		require.EqualValues(t, resp.StatusCode, v.statusCode, "idx(%d) 结果不一致", i)
		_ = resp.Body.Close()
	}
}
