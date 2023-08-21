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

// 全部都是值 Copy
func TestPointer(t *testing.T) {
	u := User{Age: 10}
	c := &u

	fmt.Println(c.Age)
	b := *c
	d := c
	b.Age = 20
	fmt.Println(c.Age)
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p", &d)

	users := []*User{
		{Age: 10},
		{Age: 20},
	}
	for i := range users {
		v := *users[i]
		v.Age = 30
	}

}

// 遍历时， v 的地址是相同的，不要对 v 做取地址操作
func TestUserPointer(t *testing.T) {
	users := make([]*User, 0, 10)
	a := User{Age: 11}
	b := User{Age: 12}
	users = append(users, &a)
	users = append(users, &b)
	fmt.Printf("\t  a:%p\t", &a)
	fmt.Printf("b:%p\n", &b)

	for i := range users {
		v := users[i]
		fmt.Printf("%d range:%p\t", i, v)
	}
	fmt.Println()

	newUsers := make([]*User, 0, 10)
	for i, v := range users {
		fmt.Printf("%d range2:%p\t", i, v)
		newUsers = append(newUsers, v)
	}

	fmt.Println()
	for _, v := range newUsers {
		fmt.Println(v.Age)
		fmt.Println(&v)
	}
}
