package channels

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	c1 := make(chan any)
	close(c1)
	c2 := make(chan any)
	close(c2)

	var c1Count, c2Count int
	for i := 0; i < 100; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count:%d\nc2Count:%d\n", c1Count, c2Count)
}

// 不可变数据是理想的，它是隐式地并行安全的。
// 每个并发进程可能对相同的数据进行操作，但不可能对其进行修改。如果要创建新数据，必须创建具有所修改的数据的新副本
// 约束可以承担更小的认知负担，约束并发指针的技术比简单传递值的副本要复杂一点
