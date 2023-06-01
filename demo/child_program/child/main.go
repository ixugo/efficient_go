package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	i := 0
	fmt.Println("123")
	for input.Scan() {
		text := input.Text()

		b, _ := json.Marshal(map[string]any{
			"text": text,
			"i":    i,
		})
		fmt.Println(string(b))
		i++
	}
	fmt.Print("end")
}
