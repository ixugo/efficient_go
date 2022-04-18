package allocationandgc

import "testing"

// 测试链表和数组遍历的速度
// go test -v array_and_list_test.go -run none -bench . -benchtime=3s -benchmem

const (
	rows    = 10000
	columns = 10000
)

var total int

// BenchmarkRowArray 访问效率更高
func BenchmarkRowArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		total = readArrayRow()
	}
}
func BenchmarkColumnArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		total = reaedArrayColumn()
	}
}

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readList()
	}
}

var array [rows][columns]int

// readArrayRow 先行后列处理
func readArrayRow() int {
	var total int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			total += array[i][j]
		}
	}
	return total
}

// reaedArrayColumn 先列后行处理
func reaedArrayColumn() int {
	var total int
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			total += array[j][i]
		}
	}
	return total
}

func readList() int {
	cur := list

	var total int
	for cur != nil {
		total += cur.v
		cur = cur.next
	}
	return total
}

type node struct {
	v    int
	next *node
}

var list = func() *node {
	header := createNode(0)

	cur := header
	for i := 0; i < rows*columns; i++ {
		cur.next = createNode(0)
		cur = cur.next
	}

	return header
}()

func createNode(v int) *node {
	return &node{v, nil}
}
