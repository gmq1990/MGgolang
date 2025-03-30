package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tTimeout() {
	// 3秒未执行完，则任务超时
	result := make(chan int)
	// timeout := make(chan int)
	// 任务例程
	go func() {
		interval := time.Duration(rand.Intn(10)) * time.Second
		fmt.Println(interval)
		time.Sleep(interval)
		result <- 0
	}()
	// 超时例程
	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	timeout <- 0
	// }()

	select {
	case <-result:
		fmt.Println("执行完成")
	case <-time.After(3 * time.Second):
		fmt.Println("执行超时")
	}
}
