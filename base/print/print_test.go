package print_test

import (
	"fmt"
	"testing"
)

type Err struct {
	Message string
}

func (e *Err) String() string {
	return "msg : " + e.Message
}

func TestPrintln(t *testing.T) {
	e := Err{Message: "test"}
	fmt.Println(e)
}

func asd(a, b, c, d int) {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
