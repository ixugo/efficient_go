package main

import (
	"fmt"
	"sync"
)

func main() {
	// 启动 2 个协程，交替打印 a,b。 a+b 总共字符 100 个

	var wg, awg, bwg sync.WaitGroup
	wg.Add(2)

	awg.Add(1)
	bwg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			bwg.Wait()
			fmt.Println("b")
			bwg.Add(1)
			awg.Done()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			awg.Wait()
			fmt.Println("a")
			awg.Add(1)
			bwg.Done()
		}
	}()
	awg.Done()
	wg.Wait()
	fmt.Println("end")
}
