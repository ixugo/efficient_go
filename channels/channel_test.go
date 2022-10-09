package channels

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 通道的三个基本模式
// waitForTask
// waitForResult
// waitForFinished

// 把任务交过去
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("recv'd signal : ", p)
	}()

	time.Sleep(500 * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------end-------------")
}

// 等协程的结果
func waitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sned signal")
	}()
	p := <-ch
	fmt.Println("manage : recv'd signal : ", p)

	time.Sleep(1 * time.Second)
	fmt.Println("-------------end-------------")
}

// 等待任务完成，
func waitForFinished() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(500 * time.Millisecond)
		close(ch)
		fmt.Println("employee : sned signal")
	}()

	_, ok := <-ch
	fmt.Println("manage : recv'd signal : ", ok)

	time.Sleep(1 * time.Second)
	fmt.Println("-------------end-------------")
}
func TestWaitForTask(t *testing.T) {
	waitForTask()
}

func TestWaitForResult(t *testing.T) {
	waitForResult()
}

func TestWaitForFinished(t *testing.T) {
	waitForFinished()
}

func pooling() {
	ch := make(chan string)
	const emps = 2
	for i := 0; i < emps; i++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd signal\n", emp)
		}(i)
	}

	const work = 10
	for i := 0; i < work; i++ {
		ch <- "paper" + strconv.Itoa(i)
		fmt.Println("manager : sent signal : ", i)
	}

	close(ch)
	fmt.Println("manage : recv'd signal end ")

	time.Sleep(1 * time.Second)
	fmt.Println("-------------end-------------")
}

func TestPooling(t *testing.T) {
	pooling()
}

// 允许任意数量的 goroutine 执行
func fanOut() {
	emps := 20
	ch := make(chan string, emps)
	for i := 0; i < emps; i++ {
		go func(i int) {
			time.Sleep(200 * time.Millisecond)
			ch <- "paper" + strconv.Itoa(i)
			fmt.Println("manager : sent signal : ", i)
		}(i)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manage : recv'd signal : ", emps)
		emps--
	}

	time.Sleep(2 * time.Second)
	fmt.Println("-------------end-------------")
}

func TestFanOut(t *testing.T) {
	fanOut()
}

// 限制同时执行的 goroutine 数量
func fanoutSemaphore() {
	emps := 20
	ch := make(chan string, emps)

	const cap = 5
	sem := make(chan struct{}, cap)

	for i := 0; i < emps; i++ {
		go func(i int) {
			sem <- struct{}{}
			{
				time.Sleep(200 * time.Millisecond)
				ch <- "paper" + strconv.Itoa(i)
				fmt.Println("manager : sent signal : ", i)
			}
			<-sem
		}(i)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manage : recv'd signal : ", emps)
		emps--
	}
}

func TestFanoutSemaphore(t *testing.T) {
	fanoutSemaphore()
}

// drop 尽快发现有问题的地方，并且防止恶化
// 抽象点就是往水杯注水，满了后，就让后来的水溢出去
// 这个模式可以用来测试性能的瓶颈
func drop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : recv'd signal : ", p)
		}
	}()

	const work = 20
	for i := 0; i < work; i++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent signal : ", i)
		default:
			fmt.Println("manager : dropped data : ", i)
		}
	}

	close(ch)
	fmt.Println("manager sent shutdown signal")
}

func TestDrop(t *testing.T) {
	drop()
}

// 注意，在这个例子中要使用 有缓冲通道
// 否则有可能会泄露，假设现在使用无缓冲通道，而任务超时，主任务结束。
// 而子任务还卡在那，等待 channel 发送过去。
func cancellation() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("manager : sent signal ")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("manage : recv'd signal : ", p)
	case t := <-tc:
		fmt.Println("manager : timedout :", t)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("-------------end-------------")
}
func TestCancellation(t *testing.T) {
	cancellation()
}

// 测试关闭 chan 后，for 循环，总是忘记
func TestCloseChan(t *testing.T) {
	ch := make(chan int, 5)
	go func() {
		ch <- 1
		time.Sleep(50 * time.Millisecond)
		ch <- 2
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("end")

}
