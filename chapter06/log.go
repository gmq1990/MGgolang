package main

import (
	"log"
	"os"
	"time"
)

func lLog() {
	// 日志输出
	path := "user.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)

	// fmt.Println(file, err)

	if err == nil {
		// 指定输出的位置
		log.SetOutput(file)
		// 加前缀
		log.SetPrefix("users:")
		// 加flag
		log.SetFlags(log.Flags() | log.Lshortfile) // 记录日志的当前文件、文件中产生日志的行
		log.Print(time.Now().Unix())
		file.Close()
	}
}
