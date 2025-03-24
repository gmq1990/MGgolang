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
		writer.WriteString("ABCDEF")
		writer.Write([]byte("123456"))
		// 刷入到文件
		writer.Flush()
	}

}
