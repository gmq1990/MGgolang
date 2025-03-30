package main

import "fmt"

func selectcase2() {
	channel01 := make(chan int, 1)
	channel02 := make(chan int, 1)
	channel01 <- 0
	channel02 <- 0

	go func() {
		<-channel01
	}()

	go func() {
		<-channel02
	}()
	// 写管道缓冲
	// 写入成功就执行case
	select {
	case channel01 <- 1:
		fmt.Println("channel01")
	case channel02 <- 1:
		fmt.Println("channel02")
	}
}
