package var_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

var a = 5

func TestVar(t *testing.T) {
	// 局部变量会覆盖全局变量作用域
	{
		a, err := strconv.Atoi("1")
		_ = err
		fmt.Println(a)
	}
	fmt.Println(a)
}

func TestVar2(t *testing.T) {
	{
		// 局域变量使用短声明赋值符时，2 个 err 是同一个对象
		var err error
		fmt.Printf("%p\n", &err)
		b, err := strconv.Atoi("a")
		fmt.Printf("%p\n", &err)
		_ = b
	}
}

func TestVar3(t *testing.T) {
	fmt.Println(addNum())
	fmt.Println(ummarshal())
}

func ummarshal() (any, error) {
	var b any
	return b, json.Unmarshal([]byte(`{"a":"b"}`), &b)
}

func addNum() (int, int) {
	i := 3
	fmt.Println(i)
	return i, add(&i)
}

func add(i *int) int {
	*i++
	return *i
}
