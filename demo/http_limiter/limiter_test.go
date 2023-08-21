package httplimiter

import (
	"io"
	"testing"
)

func initVal(w io.Writer) {
	s := "test1123"
	for i := 0; i < 10000; i++ {
		_, _ = io.WriteString(w, s)
	}
}

func TestLimiter(t *testing.T) {
	// buf := bytes.NewBuffer(nil)
	// initVal(buf)
	// for i := 0; i < 100; i++ {
	// 	start := time.Now()
	// 	r := io.LimitReader(buf, 100)
	// 	io.ReadAll(buf)
	// 	if s := time.Since(start); s > 300*time.Nanosecond {
	// 		fmt.Println(time.Since(start))
	// 	}
	// }
}
