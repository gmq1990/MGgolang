package main

import (
	"flag"
	"fmt"
)

func flagargs() {
	var host string
	var port int
	var verber bool
	var help bool
	// 绑定命令行参数与变量关系
	flag.IntVar(&port, "P", 22, "ssh port")
	flag.StringVar(&host, "H", "127.0.0.1", "ssh host")
	flag.BoolVar(&verber, "V", false, "detail log")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage: flagargs [-H 127.0.0.44] [-P 22] [-v] [args01]")
		flag.PrintDefaults()
	}
	// 解析命令行参数
	flag.Parse()

	if help {
		// go run xxx -h
		flag.Usage()
	} else {
		fmt.Println(host, port, verber)
		fmt.Println(flag.Args())
	}

}

// go run ./MGgolang/chapter04/stdpkg -h
