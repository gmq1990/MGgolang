package main

import (
	"bytes"
	"fmt"
)

func bBytes() {
	buffer := bytes.NewBufferString("abc123")

	// 内存中，buffer写
	buffer.Write([]byte("123"))
	buffer.WriteString("!@#")
	fmt.Println(buffer.String())

	// 内存中，buffer读。读过以后就没有了
	bytes := make([]byte, 2)
	buffer.Read(bytes)
	fmt.Println(string(bytes))

	// 读到指定符号后结束
	line, _ := buffer.ReadString('!')
	fmt.Println(line)

}
