package main

import (
	"bufio"
	"fmt"
	"os"
)

func osstd() {
	// os.Stdin
	// os.Stdout
	// 都是文件对象
	fmt.Println("xxx")             // xxx
	os.Stdout.Write([]byte("xxx")) // xxx

	bytes := make([]byte, 5)
	// 用户输入内容到控制台
	n, err := os.Stdin.Read(bytes)
	fmt.Println(n, err)
	fmt.Println(bytes)

	// 带缓冲的输入
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())

	fmt.Fprintf(os.Stdout, "%T", 1) // int
}
