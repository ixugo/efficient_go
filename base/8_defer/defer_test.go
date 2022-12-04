package defertest

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var i int

func recursion(i int) {
	go func() {
		i++
		defer func(i int) {
			recursion(i)
		}(i)
		fmt.Println("i++:", i)
		time.Sleep(100 * time.Millisecond)
	}()

}

func TestAsd(t *testing.T) {
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
		}
	}()
	recursion(0)
	time.Sleep(12 * time.Second)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}

// 一个要考虑到，函数底部的代码可能依赖于上面的代码
// 如果提现释放了，会导致后面的代码出现异常情况
// 所以底下的 defer 优先执行
/*
	函数栈
	-----------------------
	fmt.Println(1)
	-----------------------
		第二条 defer 语句  ^
	-----------------------
		第一条 defer 语句  ^
	-----------------------
*/
func TestDefer(t *testing.T) {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}
