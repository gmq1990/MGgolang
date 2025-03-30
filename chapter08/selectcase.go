package main

import "fmt"

func selectcase() {
	channel01 := make(chan int, 1)
	channel02 := make(chan int, 1)

	go func() {
		channel01 <- 1
	}()

	go func() {
		channel02 <- 2
	}()
	// 读管道缓冲
	// 读成功就执行case
	select {
	case <-channel01:
		fmt.Println("channel01")
	case <-channel02:
		fmt.Println("channel02")
	}
}
