package main

import (
	"fmt"
	"runtime/debug"
)

// go build -buildvcs=true .
// go version -m ./main    查看信息
func main() {
	info, ok := debug.ReadBuildInfo()
	fmt.Println(ok)
	fmt.Println(info.String())

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			fmt.Println(kv.Value)
		case "vcs.time":
			fmt.Println(kv.Value)
		case "vcs.modified":
			fmt.Println(kv.Value)
		}
	}
	fmt.Println("version", info.Main.Version)

}
