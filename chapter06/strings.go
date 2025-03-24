package main

import (
	"fmt"
	"io"
	"strings"
)

func sStrings() {
	// 直接读取内存中的流数据，不写入到硬盘
	reader := strings.NewReader("abcedf123456")

	bytes := make([]byte, 3)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		} else {
			fmt.Println(n, bytes[:n])
		}
	}

	// 给内存的流中写数据
	var builder strings.Builder
	// 写字节数组
	builder.Write([]byte("abc"))
	// 写字符串
	builder.WriteString("dfvadf123")
	// 将2者拼接
	fmt.Println(builder.String())
}
