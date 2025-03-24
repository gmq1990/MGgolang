package main

import (
	"fmt"
	"io"
	"os"
)

func readfile() {
	// 定位文件位置
	path := "user.txt"
	// 打开文件
	file, err := os.Open(path)
	fmt.Printf("%T\n", file) // *os.File
	if err != nil {
		fmt.Println(err)
	} else {
		var bytes []byte = make([]byte, 20)
		// 读取文件内容
		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Print(string(bytes[:n]))
			}
		}
		// 关闭文件
		file.Close()

		// fmt.Println(bytes) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		// n, err := file.Read(bytes)
		// fmt.Println(n, err)    // 10 <nil>
		// fmt.Println(bytes[:n]) // [97 100 102 97 100 110 118 106 51 51]
		// fmt.Println(string(bytes[:n]))

		// n, err = file.Read(bytes)
		// fmt.Println(string(bytes[:n]))

		// n, err = file.Read(bytes)
		// fmt.Println(n, err) // 0 EOF
	}

}
