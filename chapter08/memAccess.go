package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memAccess() {
	var counter int
	group := &sync.WaitGroup{}

	incr := func() {
		for i := 0; i < 100; i++ {
			counter++
			runtime.Gosched()
		}
		group.Done()
	}

	decr := func() {
		for i := 0; i < 100; i++ {
			counter--
			runtime.Gosched()
		}
		group.Done()
	}

	k := 0
	for i := 0; i < 10; i++ {
		group.Add(2)
		go incr()
		go decr()
		k++
	}

	group.Wait()
	fmt.Println(counter, k) // counter不固定，k==10
}
