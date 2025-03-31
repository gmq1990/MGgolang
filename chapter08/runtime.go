package main

import (
	"fmt"
	"runtime"
	"time"
)

func rRuntime() {
	fmt.Println(runtime.GOROOT())

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(1))

	go func() {
		time.Sleep(1 * time.Second)
		runtime.Gosched()
	}()

	fmt.Println(runtime.NumGoroutine()) // 2  一个主例程，一个工作例程
}
