package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const historyFile = "/Users/xugo/.zsh_history"
	f, err := os.Open(historyFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	cache := make(map[string]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		commandParts := strings.Split(line, ";")
		if len(commandParts) != 2 {
			fmt.Println(strconv.Quote(line))
			continue
		}
		command := strings.TrimSpace(commandParts[len(commandParts)-1])
		if _, ok := cache[command]; !ok {
			lines = append(lines, line)
			cache[command] = struct{}{}
		}
	}

	uniqueHistoryFile := "./unique_command_history.txt"
	w, err := os.Create(uniqueHistoryFile)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer w.Close()

	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}
