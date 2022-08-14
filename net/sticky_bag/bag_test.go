// 演示网络粘包/分包
// 在这个演示中，中文字符因为分包的原因导致显示乱码

package stickybag

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"
)

func server(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			fmt.Println("发现新的链接,", conn.RemoteAddr())
			defer conn.Close()

			// ===============1 =============
			// input := bufio.NewScanner(conn)
			// for input.Scan() {
			// 	s := input.Text()
			// 	fmt.Println("recv:", s)
			// }

			// ===============2 ================
			// var buf [256]byte
			// for {

			// 	// n, err := io.ReadFull(conn, buf[:])
			// 	n, err := conn.Read(buf[0:])
			// 	if err == io.EOF {
			// 		continue
			// 	} else if err != nil {
			// 		fmt.Println("read err:", err)
			// 		break
			// 	}
			// 	fmt.Println("recv:", string(buf[0:n]))
			// }

			// ==============3======================
			b := bufio.NewReader(conn)
			line, err := b.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Println(line)
		}(conn)
	}
}

func TestStickyBag(t *testing.T) {
	addr := ":" + strconv.Itoa(rand.Intn(10000)+10000)
	go server(addr)
	time.Sleep(100 * time.Millisecond)

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < 50; i++ {
		_, err = conn.Write([]byte("[一个完整的数据包]" + strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(2 * time.Second)
}
