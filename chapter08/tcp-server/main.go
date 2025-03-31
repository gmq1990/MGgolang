package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := "0.0.0.0:9999"
	// 配置连接
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// 延迟关闭服务器
	defer listener.Close()
	fmt.Printf("Listen: %s\n", addr)
	for {
		// 监听客户端连接
		client, err := listener.Accept()
		if err == nil {
			client.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
			// 关闭客户连接
			client.Close()
		}
	}
}
