package fp

import (
	"fmt"
	"testing"
	"time"
)

func Fibonacci(x int) int {
	if x == 0 {
		return 0
	}
	if x <= 2 {
		return 1
	}
	return Fibonacci(x-2) + Fibonacci(x-1)
}

func TestFibonacci(t *testing.T) {
	defer timeCost(time.Now())
	r := Fibonacci(38)
	t.Log(r)
}
func TestMemoize(t *testing.T) {
	defer useTimeFn()
	fibMem := Memoize(Fibonacci)
	r := fibMem(38)
	r = fibMem(38)
	r = fibMem(38)
	r = fibMem(38)
	t.Log(r)
}

// timeCost 耗时统计
func timeCost(start time.Time) {
	fmt.Printf("time cost: %v\n", time.Since(start))
}

var useTimeFn = UseTime()

func UseTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("time cost: %v\n", time.Since(start))
	}
}

func TestTimeCost(t *testing.T) {
	defer timeCost(time.Now())
	time.Sleep(1 * time.Second)

}

type Memoized func(int) int

var fibMem = Memoize(fib)

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func FibMemoized(n int) int {
	return fibMem(n)
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}
