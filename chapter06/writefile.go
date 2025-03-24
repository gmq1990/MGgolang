package main

import (
	"fmt"
	"os"
)

func writefile() {
	// 创建文件
	path := "user2.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		// 写文件
		file.Write([]byte("abc123!@#")) // 切片
		file.WriteString("123xyz")      // 字符串
		// 关闭文件
		file.Close()
	}
}
