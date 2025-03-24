package main

import (
	"bufio"
	"fmt"
	"os"
)

func sScanner() {
	// 带缓冲的读（扫描）
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()
		var i int
		scanner := bufio.NewScanner(file)
		// Scan(): 判断是否到末尾
		for scanner.Scan() {
			// 读取一行
			fmt.Println(i, scanner.Text())
			i++
		}
	}
}
