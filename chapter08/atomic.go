package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func aAtomic() {
	// 原子操作：
	// 使用int32/int64
	var counter int32
	group := &sync.WaitGroup{}

	incr := func() {
		defer group.Done()
		for i := 0; i < 100; i++ {
			// 将+1转为原子操作
			atomic.AddInt32(&counter, 1)
			runtime.Gosched()
		}
	}

	decr := func() {
		for i := 0; i < 100; i++ {
			// 将-1转为原子操作
			atomic.AddInt32(&counter, -1)
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
