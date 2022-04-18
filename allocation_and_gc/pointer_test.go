package allocationandgc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type user struct {
	name string
}

func (u user) Say(s string) string {
	return fmt.Sprintf("%s : %s", u.name, s)
}

type People interface {
	Say(string) string
}

type admin struct {
	user
	level int8
}

// TestValueCopy 值语义的拷贝
func TestValueCopy(t *testing.T) {
	u := user{"bob"}

	slice := []People{u, &u}
	u.name = "alice"
	require.NotEqualValues(t, slice[0].Say(""), slice[1].Say(""))

	a := admin{
		user:  u,
		level: 1,
	}
	u.name = "admin"
	require.NotEqualValues(t, a.name, "admin")
}
