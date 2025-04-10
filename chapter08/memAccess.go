package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memAccess() {
	var counter int
	group := &sync.WaitGroup{}
	lock := &sync.Mutex{} // mutual exclusion。互斥

	incr := func() {
		defer group.Done()

		for i := 0; i < 100; i++ {
			lock.Lock()   // 加锁(互斥)
			counter++     // a1.拿counter(0) a2.counter+1 a3.存counter(1)
			lock.Unlock() // 释放锁

			runtime.Gosched()
		}

	}

	decr := func() {
		for i := 0; i < 100; i++ {
			lock.Lock()
			counter-- // b1.拿counter(0) b2.counter-1 b3.counter(-1)
			lock.Unlock()

			runtime.Gosched()
		}
		group.Done()
	}

	// 并发后，可能的执行顺序：
	// a1 -> b1 -> a2 -> a3(1) -> b2 -> b3(-1)
	// a1 -> a2 -> a3(1) -> b1(1) -> b2 -> b3(0)
	// 导致每次执行结果，counter都不一样
	// 这个时候，需要给例程加锁
	k := 0
	for i := 0; i < 10; i++ {
		group.Add(2)
		go incr()
		go decr()
		k++
	}

	group.Wait()
	fmt.Println(counter, k) // 加锁前，counter不固定，k==10 // 加锁后，counter==0，k==10
}
