package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	input := bufio.NewScanner(os.Stdin)
	for {
		// 监听客户端连接
		client, err := listener.Accept()
		// 延迟关闭客户连接
		defer client.Close()

		if err == nil {
			reader := bufio.NewReader(client)
			writer := bufio.NewWriter(client)

			fmt.Printf("客户端%s已连接.\n", client.RemoteAddr())
			for {
				fmt.Print("请输入(q退出): ")
				input.Scan()
				text := input.Text()
				if text == "q" || text == "Q" {
					break
				}
				_, err := writer.WriteString(text + "\n")
				if err != nil {
					break
				}
				writer.Flush()

				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Print("客户端: ", line)
			}
			fmt.Printf("客户端%s已断开.\n", client.RemoteAddr())
		}
	}
}
