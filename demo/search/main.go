package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var ip = flag.String("ip", "192.168.2.1", "ip地址")

type Body struct {
	Version string
}

func main() {
	flag.Parse()
	ips := strings.Split(*ip, ".")
	if len(ips) != 4 {
		fmt.Println("错误的 IP")
		return
	}
	ch := make(chan string, 50)

	var wg sync.WaitGroup
	const maxG = 200
	wg.Add(maxG)
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	for i := 0; i < maxG; i++ {
		go func(i int) {
			defer wg.Done()
			for v := range ch {
				req, _ := http.NewRequest(http.MethodGet, v, nil)
				resp, err := client.Do(req)
				if err != nil {
					// fmt.Println("err", err)
					continue
				}
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					continue
				}
				var data Body
				if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
					if data.Version != "" {
						v := strings.TrimRight(v, "/healthcheck")
						fmt.Printf("NTD: %s,  version:%s\n", v, data.Version)
					}
				}
			}
		}(i)
	}

	value := make([]string, 4)
	fmt.Println("EasyNTS 搜索开始")
	now := time.Now()
	for i := 2; i <= 255; i++ {
		copy(value, ips)
		value[3] = strconv.Itoa(i)
		// ch <- fmt.Sprintf("http://%s:%s/healthcheck", strings.Join(value, "."), "80")
		// ch <- fmt.Sprintf("http://%s:%s/healthcheck", strings.Join(value, "."), "8026")
		ch <- fmt.Sprintf("http://%s:%s/healthcheck", strings.Join(value, "."), "80")
		ch <- fmt.Sprintf("http://%s:%s/healthcheck", strings.Join(value, "."), "8026")
	}
	close(ch)
	wg.Wait()
	fmt.Printf("搜索结束, 用时:%s\n", time.Since(now))
}
