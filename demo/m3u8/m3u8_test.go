package m3u8

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/grafov/m3u8"
)

func TestAppend(t *testing.T) {
	const file = `/Users/xugo/Documents/efficient_go/demo/m3u8/index.m3u8`
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(10 * time.Millisecond)
			m3u8Content, err := os.ReadFile(file)
			if err != nil {
				panic(err)
			}
			if err := os.WriteFile(file, []byte(string(m3u8Content)+`\n#EXTINF:2.933333`), os.ModePerm); err != nil {
				panic(err)
			}
		}
	}()
	go func() {
		wg.Done()
		for i := 0; i < 30; i++ {
			time.Sleep(10 * time.Millisecond)
			m3u8Content, err := os.ReadFile(file)
			if err != nil {
				panic(err)
			}
			name := filepath.Join(filepath.Dir(file), strconv.Itoa(i))
			if err := os.WriteFile(name, m3u8Content, os.ModePerm); err != nil {
				panic(err)
			}
		}
	}()
	wg.Wait()
	fmt.Println("end")
}

func TestCountM3u8(t *testing.T) {
	r, err := countM3u8Duration(`/Users/xugo/Documents/efficient_go/demo/m3u8/index.m3u8`)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func countM3u8Duration(path string) (time.Duration, error) {
	m3u8Content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading M3U8 file:", err)
		return 0, err
	}

	// 将 M3U8 内容分割成行
	lines := strings.Split(string(m3u8Content), "\n")

	// 初始化总持续时间为零
	totalDuration := time.Duration(0)

	// 遍历每行
	for _, line := range lines {
		if strings.HasPrefix(line, "#EXTINF:") {
			// 提取持续时间信息
			durationStr := strings.TrimRight(strings.TrimPrefix(line, "#EXTINF:"), ",")
			duration, err := time.ParseDuration(durationStr + "s")
			if err != nil {
				fmt.Println("Error parsing duration:", err)
				return 0, err
			}
			totalDuration += duration
		}
	}

	return totalDuration, nil
}

func TestM3u8(t *testing.T) {
	p, err := m3u8.NewMediaPlaylist(3, 10)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		if err := p.Append(fmt.Sprintf("test%d.ts", i), 3.0, ""); err != nil {
			panic(err)
		}
	}
	p.Close()
	fmt.Println(p.Encode().String())
}
