package pointer

import (
	"fmt"
	"testing"
)

type User struct {
	Age int
}

func (u User) SetAge(v int) {
	u.Age = v
}

func (u *User) SetAge2(v int) {
	u.Age = v
}

func TestUser(t *testing.T) {
	var u User
	u.SetAge(5)
	fmt.Println(u.Age)
	u.SetAge2(9)
	fmt.Println(u.Age)
}

// TestForRange
func TestForRange(t *testing.T) {
	users := []*User{
		{Age: 10},
		{Age: 11},
		{Age: 12},
	}

	fmt.Println(users)
	for _, v := range users {
		v.Age = 5
	}
	fmt.Println(users)
}

// TestForRange2
func TestForRange2(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for i, val := range slice {
		m[i] = &val
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
