package ijson

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type User struct {
	Age    int `json:"age,omitempty,string"`
	High   int `json:"-"`
	Weight int `json:"weight"`
}

func TestUser(t *testing.T) {
	s := `{ "age":"15", "weight":10, "High":5 }`
	var u User
	err := json.Unmarshal([]byte(s), &u)
	require.NoError(t, err)
	fmt.Printf("%+v\n", u)

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
