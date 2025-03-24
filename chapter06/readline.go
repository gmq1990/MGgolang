package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func rReadLine() {
	// 带缓冲的读
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			// 返回一行的切片。一行显示不了时，截断行，isPrefix会为true
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			fmt.Println(isPrefix, err, string(line))
		}
	}

}
