package file

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	const path = "data.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestPointer(t *testing.T) {
	a := 1
	b := 2
	c := 3
	arr := []*int{&a, &b, &c}
	fmt.Printf("a:%p\t", &a)
	fmt.Printf("b:%p\t", &b)
	fmt.Printf("c:%p\t\n", &c)
	for _, v := range arr {
		*v = 1
	}
	fmt.Println()

	for _, v := range arr {
		a := *v
		a = 1
		_ = a
	}
	fmt.Println(a, b, c)
	fmt.Println(arr)

}
