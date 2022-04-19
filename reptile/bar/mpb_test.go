package bar

import (
	"crypto/rand"
	"fmt"
	"io"
	"sync"
	"testing"
)

func TestNewMpb(t *testing.T) {
	p := NewMpb()
	const total = 1024 * 1024 * 20

	str := [...]string{
		"2.1 登录Linux",
		"学会使用七个基本的Linux命令行工具",
		"学会使用七x命令行工具",
		"学会使具",
		"学具",
		"学七x令行工具",
		"学会使用七x令行工具",
		"1.1 了解Linux发行版本",
		"模块一: 基本命令 基本命令",
		"课程开篇介绍",
		"1.4 安装Linux",
		"1.3 安装Ubuntu 18.04 LTS",
		"1.3 安装U 18.04 LTS",
		"1.3.04 LTS",
	}
	reader := make([]io.Reader, len(str))
	for i := range reader {
		reader[i] = io.LimitReader(rand.Reader, total)
	}

	var wg sync.WaitGroup

	for i := 0; i < len(reader); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = p.Copy(str[i], total, io.Discard, reader[i])
		}(i)
	}

	wg.Wait()
	p.Wait()
}

func TestSprintf(t *testing.T) {
	fmt.Printf("'%-5s'\n", "嗨你好")
	fmt.Printf("'% 3.1f'\n", 24.5234)
	fmt.Printf("'% 3.1f | % 3.1f'\n", 24.5234, 245.4443)
	fmt.Printf("'%.1f'\n", 234.5234)
}
