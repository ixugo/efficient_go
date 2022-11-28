package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	dir        = flag.String("dir", "", "指定目录")
	suffix     = flag.String("s", "", "指定后缀")
	replaceKey = flag.String("e", "", "指定替换字符串")
	key        = flag.String("key", "", "新的字符串")

	add = flag.String("add", "", "添加后缀")
)

func main() {
	flag.Parse()
	if *dir == "" {
		*dir = filepath.Dir(os.Args[0])
	}

	var i int
	if err := filepath.WalkDir(*dir, func(path string, d fs.DirEntry, err error) error {
		if path == *dir {
			return nil
		}
		fmt.Println(d.Name())
		i++
		if *replaceKey != "" && strings.Contains(d.Name(), *replaceKey) {
			nPath := filepath.Join(filepath.Dir(path), strings.ReplaceAll(d.Name(), *replaceKey, *key))
			return os.Rename(path, nPath)
		}
		// 一般文件夹可能并不想添加后缀
		// 所以将 后缀替换 和 添加后缀拆分开
		if *suffix != "" && strings.HasSuffix(d.Name(), *suffix) {
			return os.Rename(path, strings.TrimRight(path, *suffix)+*key)
		}
		if *add != "" {
			return os.Rename(path, path+*add)
		}
		i--
		return nil
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("操作完成，共处理文件: %d 个\n", i)
}
