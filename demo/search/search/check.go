package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// checkSD 检查是否赛达
func checkSD(addr string) {
	client := http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/user/login", addr), nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	b, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(b), "xsw.") {
		return
	}
	fmt.Println("赛达设备: ", addr)

	// b, _ := io.ReadAll(resp.Body)
	// if strings.Contains(string(b),"GB28181")
	// var data Body
	// if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
	// 	if data.Version != "" {
	// 		v := strings.TrimRight(v, "/healthcheck")
	// 		fmt.Printf("NTD: %s,  version:%s\n", v, data.Version)
	// 	}
	// }
}

// checkHK 检查是否海康
func checkHK(addr string) {
	client := http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/doc/i18n/Languages.json", addr), nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	b, _ := io.ReadAll(resp.Body)
	if strings.HasPrefix(string(b), `{"Languages"`) {
		fmt.Println("海康设备:", addr)
	}
}

func checkHK2(addr string) {
	client := http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/doc/i18n/zh/Login.json", addr), nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	// b, _ := io.ReadAll(resp.Body)
	// if strings.HasPrefix(string(b), `{"Languages"`) {
	fmt.Println("海康设备:", addr)
	// }
}

// checkNTD 检查是否 ntd
func checkNTD(addr string) {
	client := http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/healthcheck", addr), nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	b, _ := io.ReadAll(resp.Body)
	if strings.Contains(string(b), `easyntd`) {
		fmt.Println("EasyNTD:", addr)
	}
}

// 检查电信设备
func checkDX(addr string) {
	client := http.Client{
		Timeout: 2500 * time.Millisecond,
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s", addr), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	b, _ := io.ReadAll(resp.Body)
	if strings.Contains(string(b), `upper-right`) {
		fmt.Println("中国电信:", addr)
	}
}

func checkYuShi(addr string) {
	client := http.Client{
		Timeout: 2500 * time.Millisecond,
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/LAPI/V1.0/System/Security/RSA", addr), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	b, _ := io.ReadAll(resp.Body)
	if strings.Contains(string(b), `"/LAPI`) {
		fmt.Println("宇视:", addr)
	}
}

func check(ch chan *Value) {
	for v := range ch {
		switch v.Atype {
		case "hk":
			// checkHK(v.Addr)
			checkHK2(v.Addr)
		case "sd":
			checkSD(v.Addr)
		case "ntd":
			checkNTD(v.Addr)
		case "dx":
			checkDX(v.Addr)
		case "ys":
			checkYuShi(v.Addr)
		default:
			// checkHK(v.Addr)
			checkHK2(v.Addr)
			checkSD(v.Addr)
			checkNTD(v.Addr)
			checkYuShi(v.Addr)
		}
	}
}
