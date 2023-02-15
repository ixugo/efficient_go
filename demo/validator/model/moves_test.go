package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMoves(t *testing.T) {
	var m Movie
	m.Runtime = 12
	m.Title = "test"

	b, _ := json.Marshal(m)
	fmt.Println(string(b))

	runtime := `{"runtime":"15 mins}`
	if err := json.Unmarshal([]byte(runtime), &m); err != nil {
		t.Fatal(err)
	}
	fmt.Println(m.Runtime)
}

type A struct {
	a int
	d map[string]string
}

func (a *A) clone() *A {
	l := *a
	return &l
}

func TestASD(t *testing.T) {
	a := A{}
	b := &a
	c := *b
	fmt.Printf("a:%p\nb:%p\nc:%p\n", &a, b, &c)
	fmt.Printf("%p\n%p", &a, a.clone())

}

func TestASDV(t *testing.T) {
	a := A{d: make(map[string]string)}
	a.d["123"] = "123"
	fmt.Printf("%p\n", &a.d)
	b := a
	fmt.Printf("%p\n", &b.d)
	fmt.Println(b.d["123"])

}
