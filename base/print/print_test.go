package print_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

func TestIO(t *testing.T) {
	buf := bytes.NewReader([]byte("Hello world"))
	reader := io.LimitReader(buf, 5)

	w := io.MultiWriter(os.Stdout, os.Stderr)
	read := io.TeeReader(reader, w)
	b, _ := io.ReadAll(read)
	fmt.Println(string(b))

	// reads := io.MultiReader(buf, read))

}
