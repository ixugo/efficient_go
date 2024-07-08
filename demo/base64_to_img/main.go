package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var baseStr = flag.String("base64", "", "base64 值")
var baseFile = flag.String("file", "", "base64 file 路径")

func main() {
	flag.Parse()

	var sss string
	if *baseStr != "" {
		sss = *baseStr
	} else if *baseFile != "" {
		b, err := os.ReadFile(*baseFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		sss = string(b)
	} else {
		fmt.Println("参数有误")
	}

	sss = strings.ReplaceAll(sss, `\/`, "/")
	b, err := base64.StdEncoding.DecodeString(sss)
	if err != nil {
		fmt.Println(">>>", err)
		return
	}
	path := filepath.Join(filepath.Dir(os.Args[0]), "out.jpeg")
	_ = os.WriteFile(path, b, os.ModePerm)

	fmt.Println("\n\n图片已写入 >>", path)
}
