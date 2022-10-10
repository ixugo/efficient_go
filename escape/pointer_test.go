package escape

import (
	"fmt"
	"reflect"
	"testing"
)

// go test -gcflags "-m" ./pointer_test.go
func TestEscapePointer(t *testing.T) {
	escapePointer()
}

// escapePointer
func escapePointer() {
	a := 5
	reflect.TypeOf(a).Kind()
	b := 9
	reflect.ValueOf(b)
	c := 10
	fmt.Println(c)
}
