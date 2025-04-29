package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch
			if !ok {
				return
			}
			fmt.Println("a", v)
			v++
			if v > 100 {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	// 1,3,5,7,9
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch
			if !ok {
				return
			}
			fmt.Println("b", v)
			v++
			if v > 100 {
				close(ch)
				return
			}
			ch <- v
		}
	}()

	ch <- 1
	wg.Wait()
	fmt.Println("end")
}
