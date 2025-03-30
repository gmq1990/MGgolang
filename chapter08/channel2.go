package main

import (
	"fmt"
	"runtime"
)

func PrintChars3(n int, channel chan int) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%d: %c\n", n, ch)
		runtime.Gosched()
	}
	channel <- n
	fmt.Println()
}

func Channel2() {
	var channel chan int = make(chan int)
	// 使用chan时，要有goroutine配合使用，否则会死锁。
	for i := 0; i < 10; i++ {
		go PrintChars3(i, channel)
	}

	fmt.Println("start")
	for i := 0; i < 10; i++ {
		fmt.Println("读取完:", <-channel)
	}
	fmt.Println("over")
}
