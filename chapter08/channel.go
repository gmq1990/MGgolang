package main

import "fmt"

func Channel() {
	var notice chan bool = make(chan bool)
	fmt.Printf("%T %v\n", notice, notice) // chan string <nil>

	// 使用chan时，要有goroutine配合使用，否则会死锁。
	go func() {
		fmt.Println("go start a")
		notice <- true
		fmt.Println("go end a")
	}()

	fmt.Println("start")
	fmt.Println(<-notice)
	fmt.Println("end")
}
