package file

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const data = `0123456789ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz`

func genrateID(prefix string, length int) string {
	dataLen := len(data)
	arr := make([]byte, length)
	_, _ = rand.Read(arr)
	var buf strings.Builder
	buf.WriteString(prefix)
	for _, v := range arr {
		idx := v % uint8(dataLen)
		buf.WriteByte(data[idx])
	}
	cache := buf.String()
	return validID(cache)
}

func validID(id string) string {
	num, _ := strconv.ParseUint(hex.EncodeToString([]byte(id)), 16, 64)
	return fmt.Sprintf("%s%c", id, data[num%53])
}

func TestID(t *testing.T) {
	for range 100 {
		s := genrateID("", 8)
		ss := s[0 : len(s)-1]
		if validID(ss) != s {
			panic(ss)
		}
		fmt.Println(s)
	}
}
