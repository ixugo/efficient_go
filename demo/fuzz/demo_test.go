package fuzz

import (
	"testing"
)

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// func FuzzDemo(f *testing.F) {
// 	cases := []string{"Hello, world", " ", "!12345", "哈喽"}
// 	for _, v := range cases {
// 		f.Add(v)
// 	}
// 	f.Fuzz(func(t *testing.T, o string) {
// 		t.Log(o)
// 		rev := Reverse(Reverse(o))
// 		require.Equal(t, rev, o)
// 	})
// }

func FuzzZero(f *testing.F) {
	f.Add(512, "asd")
	f.Add(-123, "")
	f.Fuzz(func(t *testing.T, a int, v string) {
		if a < 0 && v == "" {
			t.Fatal(v)
		}
	})
}
