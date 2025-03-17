package main

import (
	"fmt"
	"os"
)

func args() {
	if len(os.Args) < 3 {
		// 命令行参数，编辑命令行输出
		fmt.Println("not enough args")
		return
	}
	fmt.Println(os.Args)
}
