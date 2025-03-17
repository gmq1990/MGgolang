package main

import "log"

func lLog() {
	log.Printf("我是Printf日志：%s", "x")
	log.Printf("我是Printf日志：%s", "x")

	// 使后面的日志带有前缀
	log.SetPrefix("prefix:")
	// 使后面的日志带有 日志文件 的flag
	log.SetFlags(log.Flags() | log.Lshortfile)

	// 打印日志，并触发panic
	// log.Panicf("我是Panic日志：%s", "x")

	log.Fatalf("我是Fatal日志：%s", "x")
	// fatal使程序退出，下面的不执行
	log.Printf("我是Printf日志：%s", "x")

}
