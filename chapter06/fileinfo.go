package main

import (
	"fmt"
	"os"
	"strings"
)

func fileinfo() {
	// 文件是否存在
	file, err := os.Open("xxxx")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
	} else {
		file.Close()
	}

	// 路径是否存在
	for _, path := range []string{"xxx", "user.txt", "MGgolang"} {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(strings.Repeat("*", 20))
			fmt.Println(fileInfo.Name())
			fmt.Println(fileInfo.IsDir())
			fmt.Println(fileInfo.Size())
			fmt.Println(fileInfo.ModTime())

			if fileInfo.IsDir() {
				dirFile, err := os.Open(path)
				if err == nil {
					defer dirFile.Close()
					// 获取所有子路径
					children, _ := dirFile.Readdir(-1)
					for _, child := range children {
						fmt.Println(child.Name(), child.IsDir(), child.Size(), child.ModTime())
					}

					// 获取所有子路径名
					names, _ := dirFile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println(name)
					}
				}
			}
		}
	}
}
