package main

import (
	"bufio"
	"os"
)

func wWriter() {
	// 带缓冲io的写

	file, err := os.Create("user.txt")
	if err == nil {
		defer file.Close()

		writer := bufio.NewWriter(file)
		// 写入字符串
		writer.WriteString("ABCDEF")
		// 写入字符数组
		writer.Write([]byte("123456"))
		// 刷入到文件
		writer.Flush()
	}

}
