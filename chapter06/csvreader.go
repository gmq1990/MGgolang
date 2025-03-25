package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func csvreader() {
	// 读取csv文件的内容
	file, err := os.Open("user.csv")
	if err == nil {
		defer file.Close()

		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(line)
			}
		}

	}
}
