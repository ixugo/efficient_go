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
