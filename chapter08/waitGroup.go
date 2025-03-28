package main

import (
	"fmt"
	"sync"
)

func PrintChars2(n string, group *sync.WaitGroup) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s: %c\n", n, ch)
	}
	fmt.Println()
	// 计数信号量 减1
	group.Done()
}

func waitGroup() {
	// 主例程等待工作例程结束后，再结束
	group := &sync.WaitGroup{}
	// 计数信号量 WaitGroup
	group.Add(3)
	go PrintChars2("1", group)
	go PrintChars2("2", group)

	PrintChars2("3", group)
	// 等待计数信号量归0
	group.Wait()
}
