package main

import (
	"fmt"
	"io"
	"os"
)

func readall() {
	// 定位文件位置
	path := "user.txt"
	// 打开文件
	file, err := os.Open(path)
	fmt.Printf("%T\n", file) // *os.File
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		// 读取全部内容
		bytes, err := io.ReadAll(file)
		fmt.Println(string(bytes), err)
	}
}
