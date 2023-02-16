package base

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
)

func TestMath1(t *testing.T) {
	// a<=b，则 a&b == a
	a := 1
	b := 15

	if a > b {
		a = b
	}
	t.Log(a & b)

	// 其它写法
	c := a
	if a > b {
		c = b
	}
	t.Log(c)
}

func TestError(t *testing.T) {
	b, err := json.Marshal(io.EOF)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}
