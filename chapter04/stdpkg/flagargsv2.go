package main

import (
	"flag"
	"fmt"
)

func flagargsv2() {
	// 绑定命令行参数与变量关系
	port := flag.Int("P", 22, "ssh port")
	host := flag.String("H", "127.0.0.1", "ssh host")
	verbor := flag.Bool("V", false, "detail log")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage: flagargs [-H 127.0.0.44] [-P 22] [-v] [args01]")
		flag.PrintDefaults()
	}
	// 解析命令行参数
	flag.Parse()
	fmt.Printf("%T %T %T %T\n", port, host, verbor, help)
	fmt.Printf("%v %v %v %v\n", *port, *host, *verbor, *help)
}

// go run ./MGgolang/chapter04/stdpkg -h
