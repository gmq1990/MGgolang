package main

import (
	"fmt"
	"os"
)

func sysdir() {
	// 缓存路径
	fmt.Println(os.TempDir())
	// 用户cache路径
	fmt.Println(os.UserCacheDir())
	// 用户主目录
	fmt.Println(os.UserHomeDir())
	// 当前运行程序的路径
	fmt.Println(os.Getwd())
}
