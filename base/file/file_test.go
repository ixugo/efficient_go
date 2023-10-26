package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFile(t *testing.T) {
	dir, _ := os.Getwd()
	target := filepath.Join(dir, "test.txt")
	file, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	go func() {
		for {
			time.Sleep(time.Second)
			b, err := os.ReadFile(target)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		}
	}()

	for i := 0; i < 1000; i++ {
		for _, v := range "Hello world\n" {
			_, err := file.WriteString(string(v))
			if err != nil {
				panic(err)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}
}
