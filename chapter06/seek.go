package main

import (
	"fmt"
	"os"
)

func sSeek() {
	// 定位文件中内容的位置
	file, _ := os.Open("text.txt")

	defer file.Close()

	bytes := make([]byte, 100)
	n, _ := file.Read(bytes)

	// 偏移量、相对位置
	// 文件开始： 0 os.SEEK_SET
	// 当前位置： 1 os.SEEK_CUR
	// 文件末尾： 2 os.SEEK_END
	fmt.Println(file.Seek(0, 1))

	fmt.Println(n, bytes[:n])
	n, err := file.Read(bytes)

	fmt.Println(n, err, bytes[:n])

}
