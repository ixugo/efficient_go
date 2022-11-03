package ip

import (
	"io"
	"net"
	"net/http"
	"time"
)

// LocalIP 获取本地IP地址
func LocalIP() string {
	conn, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
	if err != nil {
		return ""
	}
	host, _, _ := net.SplitHostPort(conn.LocalAddr().(*net.UDPAddr).String())
	if host == "" {
		return localIP()
	}
	return host
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

// ExternalIP 获取公网 IP
func ExternalIP() (string, error) {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := c.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	ip, err := io.ReadAll(resp.Body)
	return string(ip), err
}
