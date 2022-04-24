package test

import (
	"testing"
)

func TestSubTest(t *testing.T) {

	data := []struct {
		name   string
		result bool
	}{
		{"123", true},
		{"246", false},
	}

	for _, v := range data {
		tf := func(t *testing.T) {
			t.Parallel()
			t.Log(v)
		}
		t.Run(v.name, tf)
	}
}
