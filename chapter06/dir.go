package main

import (
	"fmt"
	"os"
)

func dir() {
	// os.Mkdir("test01", 0644)
	// os.Rename("test01", "test02")
	// os.Remove("test02")
	// fmt.Println(os.MkdirAll("test01/xxxx", 0777)) // 创建目录及其子目录
	fmt.Println(os.RemoveAll("test01")) // 删除目录及其子目录
}
