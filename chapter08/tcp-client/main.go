package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:9999"
	// 客户端配置连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	bytes := make([]byte, 1024)
	// 读出接到的数据
	if n, err := conn.Read(bytes); err == nil {
		fmt.Println(string(bytes[:n]))
	} else {
		fmt.Println(err)
	}
	// 客户端关闭连接
	conn.Close()

}
