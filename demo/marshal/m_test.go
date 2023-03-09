package marshall

import (
	"encoding/json"
	"fmt"
	"testing"
)

var _ json.Marshaler = new(User)

type User struct {
	Hand
	Title string
}

type Hand struct {
	Age int
}

func (h Hand) MarshalJSON() ([]byte, error) {
	type hand Hand
	return json.Marshal(hand(h))
}

func TestMarshal(t *testing.T) {
	u := User{Title: "H", Hand: Hand{Age: 10}}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
