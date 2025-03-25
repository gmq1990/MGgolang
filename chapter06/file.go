package main

import "os"

func file() {
	// 重命名文件
	os.Rename("user.log", "user.v2.log")
	// 删除文件
	os.Remove("user.txt")
}
