package gc

import "testing"

// 测试链表和数组遍历的速度
// go test -v array_and_list_test.go -run none -bench . -benchtime=3s -benchmem

const (
	rows    = 10000
	columns = 10000
)

// BenchmarkRowArray 访问效率更高
func BenchmarkRowArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeArrayRow()
	}
}
func BenchmarkColumnArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeArrayColumn()
	}
}

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeList()
	}
}

var array [rows][columns]int

// writeArrayRow 先行后列处理
func writeArrayRow() {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			array[i][j] = 1
		}
	}
}

// writeArrayColumn 先列后行处理
func writeArrayColumn() {
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			array[j][i] = 1
		}
	}
}

func writeList() {
	cur := list
	for cur != nil {
		cur.v = 1
		cur = cur.next
	}
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
