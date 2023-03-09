package ip

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLocalIP(t *testing.T) {
	conn, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
	t.Fatal("!23")
	if err != nil {
		t.Fatal(err)

	}
	host, _, _ := net.SplitHostPort(conn.LocalAddr().(*net.UDPAddr).String())
	fmt.Println(host)
}

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}

func TestPanic(t *testing.T) {
	go func() {
		time.Sleep(5 * time.Second)
	}()
	go func() {
		time.Sleep(5 * time.Second)
	}()
	go func() {
		time.Sleep(5 * time.Second)
	}()
	go func() {
		time.Sleep(5 * time.Second)
	}()
	go func() {
		time.Sleep(5 * time.Second)
		panic(nil)
	}()

	time.Sleep(10 * time.Second)
}

type ASD struct {
	A int
	B int
	C int
}

func TestASD(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var s ASD
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			defer wg.Done()
			s.A = 1
		}()
		go func() {
			defer wg.Done()
			s.B = 2
		}()
		go func() {
			wg.Done()
			s.C = 3
		}()
		wg.Wait()
	}
}

type User struct {
	Likes []int  `json:"likes"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	B     []byte
}

func TestMarshall(t *testing.T) {
	var u User
	u.B = append(u.B, []byte{'a', 'b'}...)
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}

func TestParseIP(t *testing.T) {
	testCases := []struct {
		desc     string
		expected bool
	}{
		{
			desc:     "",
			expected: false,
		},
		{
			desc:     "129.168.4.1",
			expected: true,
		},
		{
			desc:     "129.168.23.141",
			expected: true,
		},
		{
			desc:     "129.168.993.141",
			expected: false,
		},
		{
			desc:     "129.168.993.141:412",
			expected: false,
		},
		{
			desc:     "129.168.93.141:412",
			expected: false,
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			require.EqualValues(t, net.ParseIP(v.desc) != nil, v.expected)
		})
	}
}
