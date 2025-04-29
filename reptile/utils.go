package reptile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// writeFile 序列化并写入文件
func writeFile(t any, filename string) error {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, b, os.ModePerm)
}

// writeBigFile 写大文件时，避免全部读取到内存
func writeBigFile(name string, src io.Reader) error {
	file, err := os.Create(name + ".tmp")
	if err != nil {
		return err
	}
	// 优化: 避免大文件读取到内存(如视频)
	if _, err = io.Copy(file, src); err != nil {
		return err
	}
	return os.Rename(name+".tmp", name)
}

// newRequest 创建请求
func newRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "https://servicewechat.com/wx99a836acd3f12e4d/28/page-frame.html")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.42(0x18002a32) NetType/4G Language/zh_CN")
	// req.Header.Set("Accept-Encoding", "gzip,compress,br,deflate")
	req.Header.Set("Host", "cloud.cmpreading.com")
	return req, nil
}

var cli = http.Client{
	Timeout: 5 * time.Second,
}

// do 发起网络请求
func do(req *http.Request) ([]byte, error) {
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return b, err
}
