package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func rReadLineV2() {
	// 带缓冲的读
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			// 返回一段切片, 以换行符为截断点
			// line, err := reader.ReadSlice('\n')
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			// fmt.Println(err, string(line))
			fmt.Println(err, line)
		}
	}

}
