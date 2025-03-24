package main

import (
	"os"
	"strconv"
	"time"
)

func appendfile() {
	// 追加文件
	path := "user.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)

	// fmt.Println(file, err)

	if err == nil {
		// 不清空文件原有的内容，追加新内容到文件中
		file.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		file.WriteString("\n")
		file.Close()
	}

}
