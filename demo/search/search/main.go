package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	"time"
)

var ip = flag.String("ip", LocalIP(), "请输入当前的 ip 地址")
var atype = flag.String("type", "all", "请输入搜索类型 sd:赛达,hk:海康,ntd:内网穿透,all:全部,dx:电信,ys:宇视")

// LocalIP 获取本地IP地址
func LocalIP() string {
	conn, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
	if err != nil {
		return ""
	}
	host, _, _ := net.SplitHostPort(conn.LocalAddr().(*net.UDPAddr).String())
	if host != "" {
		return host
	}
	iip := strings.Split(localIP()+"/", "/")
	if len(iip) >= 2 {
		return iip[0]
	}
	return ""
}

// localIP 获取本地 IP，遇到虚拟 IP 不准确
func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	ip := ""
	for _, v := range addrs {
		net, ok := v.(*net.IPNet)
		if !ok {
			continue
		}
		if net.IP.IsMulticast() || net.IP.IsLoopback() || net.IP.IsLinkLocalMulticast() || net.IP.IsLinkLocalUnicast() {
			continue
		}
		if net.IP.To4() == nil {
			continue
		}

		ip = v.String()
	}
	return ip
}

type Value struct {
	Addr  string
	Atype string
}

var receCh = make(chan string, 20)

func main() {
	flag.Parse()
	fmt.Println("==========================================")
	fmt.Println("-------------- 超蔷搜索工具 --------------")
	fmt.Println("==========================================")
	go func() {

	}()

	// 参数检查
	ips := strings.Split(*ip, ".")
	if len(ips) != 4 {
		fmt.Println("错误的 IP")
		return
	}
	var wg sync.WaitGroup
	ch := make(chan *Value, 200)

	const maxNum = 50
	wg.Add(maxNum)
	for i := 1; i <= maxNum; i++ {
		go func() {
			defer wg.Done()
			check(ch)
		}()
	}

	fmt.Println("搜索开始 >>>", *atype, ",", *ip)
	now := time.Now()

	value := make([]string, 4)
	for i := 2; i <= 255; i++ {
		copy(value, ips)
		value[3] = strconv.Itoa(i)
		if *atype == "dx" {
			ch <- &Value{
				Addr:  fmt.Sprintf("%s:%s", strings.Join(value, "."), "8000"),
				Atype: *atype,
			}
			ch <- &Value{
				Addr:  fmt.Sprintf("%s:%s", strings.Join(value, "."), "80"),
				Atype: *atype,
			}
		} else {
			ch <- &Value{
				Addr:  fmt.Sprintf("%s:%s", strings.Join(value, "."), "80"),
				Atype: *atype,
			}
			if *atype == "ntd" {
				ch <- &Value{
					Addr:  fmt.Sprintf("%s:%s", strings.Join(value, "."), "8026"),
					Atype: *atype,
				}
			}
		}
	}
	close(ch)
	wg.Wait()
	fmt.Printf("搜索结束, 用时:%s\n", time.Since(now))

	fmt.Println("\n\n\t\t\t按回车键退出")
	_, _ = fmt.Scanln()
}
