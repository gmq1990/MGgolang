package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func filePath() {
	// 当前运行时所处的绝对路径。
	fmt.Println(filepath.Abs(".")) // /Users/{$Username}/Documents/go-study

	fmt.Println(os.Args) // [./bin/chapter06]
	// 运行文件的绝对路径
	path, _ := filepath.Abs(os.Args[0])
	// 运行文件的父目录
	dirPath := filepath.Dir(path)

	// 获取路径名/文件名。Base()
	fmt.Println(filepath.Base("./user.txt")) // user.txt
	fmt.Println(filepath.Base("/opt"))       // opt
	fmt.Println(filepath.Base(path))         // chapter06
	// 获取路径/文件的父目录，Dir()
	fmt.Println(filepath.Dir("./user.txt")) // .
	fmt.Println(filepath.Dir("/opt"))       // /
	fmt.Println(filepath.Dir(path))         // /Users/{$Username}/Documents/go-study/bin

	// 路径扩展名/文件扩展名。Ext()
	fmt.Println(filepath.Ext("./user.txt")) // .txt
	fmt.Println(filepath.Ext("/opt"))       // （空）
	fmt.Println(filepath.Ext(path))         //  (空)

	// 手动写入分割符“/”
	iniPath := dirPath + "/conf/ip.ini"
	fmt.Println(iniPath)
	// 避免路径中的分割符的影响。Join()
	fmt.Println(filepath.Join(dirPath, "conf", "ip.ini"))

	// 正则表达，获取目标目录/文件(所有)
	fmt.Println(filepath.Glob("./[u]*.txt")) //  [user.txt user2.txt] <nil>

	// 遍历所有目录
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path, info)
		return nil
	})

}
