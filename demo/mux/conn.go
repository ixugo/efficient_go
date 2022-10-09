package mux

import (
	"net"
)

type PortConn struct {
	net.Conn        // 网络连接
	resource []byte // 资源
	readMore bool
	start    int // 数组读到的位置
}

func newPortConn(conn net.Conn, resource []byte, readMore bool) *PortConn {
	return &PortConn{
		Conn:     conn,
		resource: resource,
		readMore: readMore,
	}
}

func (c *PortConn) Read(b []byte) (int, error) {
	// 传入切片小于数据，则读多少内容
	if l := len(b); l < len(c.resource)-c.start {
		c.start += l
		return copy(b, c.resource), nil
	}

	var n int
	// 当前资源未读完，全部读完
	if c.start < len(c.resource) {
		c.start = len(c.resource)
		n = copy(b, c.resource[c.start:])
		if !c.readMore {
			return n, nil
		}
	}

	// 想要读到更多数据
	n2, err := c.Conn.Read(b[n:])
	n = n + n2
	return n, err
}
