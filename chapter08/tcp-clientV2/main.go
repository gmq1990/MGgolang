package main

import (
	"bufio"
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
	// 客户端关闭连接
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	input := bufio.NewScanner(os.Stdin)

	for {
		// 读出接到的数据
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("服务器: ", line)
		fmt.Print("请输入(q退出): ")
		// 标准输入
		input.Scan()
		if input.Text() == "q" || input.Text() == "Q" {
			break
		}
		// 写数据给服务端
		_, err = writer.WriteString(input.Text() + "\n")
		if err != nil {
			break
		}
		writer.Flush()
	}
}
