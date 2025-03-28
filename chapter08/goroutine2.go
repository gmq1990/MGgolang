package main

import (
	"fmt"
	"sync"
)

func goroutine2() {
	// 主例程等待工作例程结束后，再结束
	group := &sync.WaitGroup{}
	n := 5

	// 计数信号量 WaitGroup
	group.Add(n)
	for i := 1; i <= n; i++ {
		// 闭包
		go func() {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d: %c\n", i, ch)
			}
			fmt.Println()
			// 计数信号量 减1
			group.Done()
		}()
	}
	// 等待计数信号量归0
	group.Wait()
}
