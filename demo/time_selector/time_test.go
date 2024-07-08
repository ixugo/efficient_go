package timeselector

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	list := [24]int{22: 1, 23: 0}
	// list := [24]int{22: 1, 23: 1}

	fmt.Println("start")
	var start bool
	var str string
	for i, v := range list {
		s := i
		if v == 1 && i == 23 {
			s = 0
		}

		if !start && v == 1 {
			start = true
			str = fmt.Sprintf("%02d:00", i)
		}

		if start && (s == 0 || i == len(list)-1) {
			start = false
			str += fmt.Sprintf("~%02d:00", s)
			fmt.Println(str)
		}
	}

}
