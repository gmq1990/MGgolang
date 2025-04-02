package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	dirPath         string
	defaultFileName string
	file404         string
	file500         string
)

func handle(conn net.Conn) {
	// defer必须在一个函数内，而不能在函数的某个代码块内
	defer func() {
		conn.Close()
		fmt.Println("Client Closed: ", conn.RemoteAddr())
	}()
	fmt.Println("sleep start")
	time.Sleep(time.Second * 10)
	fmt.Println("sleep end")

	// 处理客户端数据
	reader := bufio.NewReader(conn)
	// 读取数据
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		// 正常处理
		fmt.Println(line)
		// fmt.Printf("%q\n", strings.Fields(line))
		nodes := strings.Fields(line)
		// 获取客户端请求的逻辑路径
		path := filepath.Join(dirPath, nodes[1])
		if info, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				path = file404
			} else {
				path = file500
			}
		} else {
			// 目录
			if info.IsDir() {
				path = filepath.Join(path, defaultFileName)
			}
			// 文件

		}

		// 再次对path进行检查（404.html/505.html/index/html可能也不存在）
		if _, err := os.Stat(path); err == nil {
			// conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n")))
			// bufio
			// fmt.Fprint, fmt.Fprintf
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Server: kkserver1.0\r\n")
			fmt.Fprint(conn, "\r\n")
			bytes, _ := os.ReadFile(path)
			conn.Write(bytes)
		} else {
			fmt.Fprint(conn, "HTTP/1.1 404 NotFound\r\n")
		}
	}
}

func init() {
	// 进程（运行文件）的绝对路径
	binPath, _ := filepath.Abs(os.Args[0])
	// 进程的父路径
	dirPath := filepath.Dir(binPath)

	defaultFileName = "index.html"
	file404 = filepath.Join(dirPath, "404.html")
	file500 = filepath.Join(dirPath, "500.html")
}

func httpserver() {

	addr := ":9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer listener.Close()
	fmt.Println("Listening on: ", listener.Addr())

	for {
		// 接收客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Client Connected: ", conn.RemoteAddr())

		go handle(conn)

	}
}
