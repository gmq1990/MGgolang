package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintChars(n string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s: %c\n", n, ch)
		runtime.Gosched() // 执行一会，让出cpu。由另外2个例程去抢cpu。
	}
	fmt.Println()
}

func print() {
	go PrintChars("1")
	go PrintChars("2")

	PrintChars("3")
	time.Sleep(3 * time.Second)
}
