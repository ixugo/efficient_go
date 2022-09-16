package system

import (
	"net"
)

// LocalIP 获取本地 IP
func LocalIP() string {
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

		if net.IP.IsLoopback() {
			continue
		}

		if net.IP.To4() == nil {
			continue
		}

		ip = v.String()
	}
	return ip

}
