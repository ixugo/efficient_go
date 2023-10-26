package file

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	const path = "data.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
