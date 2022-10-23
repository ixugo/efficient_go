// Package channels 如何构建稳定的 channel?

/*
-------------------------------------------------------
-		|	status				|	result
-------------------------------------------------------
read	|	nil					|	阻塞
read	|	open and not full	|	读取成功
read	|	open and empty		|	阻塞
read	|	closed				|	读取零值
read	|	only write			|	编译错误
-------------------------------------------------------
write	|	open and full		|	阻塞
write	|	open and not full	|	写入
write	|	closed				|	panic
write	|	only read			|	compile error
-------------------------------------------------------
close	|	nil					|	panic
close	|	open and not full	|	关闭，读 channel 成功
close	|	open and empty	    |	关闭，读 channel 为零值
close	| 	closed		     	|	panic
close	|	only read			|	编译错误
-------------------------------------------------------
*/

/*

从上面的图表中，可以看到:
+ 有三种操作会导致阻塞
+ 有三种操作会导致 panic

如何写出健壮和稳定的程序呢?
*/

package channels

import (
	"fmt"
	"testing"
	"time"
)

// chanOwner ...
// 拥有 channel 的 goroutine 应该具备如下要素:
// 1. 实例化
// 2. 执行写操作，或将所有权传递给另外一个 goroutine
// 3. 控制关闭
// 4. 将读取权暴露出来
func chanOwner() <-chan int {
	stream := make(chan int, 3)
	go func() {
		// channel 的所有者处理 发送和关闭
		defer close(stream)
		for i := 0; i <= 5; i++ {
			stream <- i
		}
	}()
	return stream
}

func TestChanOwner(t *testing.T) {
	stream := chanOwner()
	// 消费者只需要知道如何处理阻塞读取和 channel 关闭的值
	for v := range stream {
		fmt.Print(v)
	}
	fmt.Println("end")
}

// 约定: 如果 goroutine 负责创建 goroutine，也必须负责确保它可以停止。
func TestDowork(t *testing.T) {
	doWork := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer fmt.Println("dowork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return

				}
			}
		}()
		return terminated
	}

	done := make(chan struct{})
	terminated := doWork(done, nil)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("canceling doWork gouroutine...")
		close(done)
	}()
	<-terminated
}
