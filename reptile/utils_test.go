package reptile

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodPost, "http://baidu.com", strings.NewReader(""),
	)
	require.NoError(t, err)
	// 新创建的 header 默认为空
	require.EqualValues(t, len(req.Header), 0)
}

// TestWriteBigFile 测试写入大文件
func TestWriteBigFile(t *testing.T) {
	reader := strings.NewReader("Hello")
	err := writeBigFile("s.txt", reader)
	require.NoError(t, err)
}

// BenchmarkRead copy 比 readall 更节省内存
func BenchmarkRead(b *testing.B) {
	// ReadAll 将数据全部读到缓冲区
	b.Run("ReadAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reader := strings.NewReader("Hello")
			s := bytes.NewBuffer(nil)
			b, _ := io.ReadAll(reader)
			s.Write(b)
		}
	})

	// copy 使用固定长度的 buffer 作为缓冲区，将 src 写入 dst
	b.Run("Copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reader := strings.NewReader("Hello")
			s := bytes.NewBuffer(nil)
			_, _ = io.Copy(s, reader)
		}
	})
}
