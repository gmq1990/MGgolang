package main

import (
	"fmt"
	"sync"
)

// 自定义一个对象池

type New func() interface{}

type Pool struct {
	mutex   sync.Mutex
	objects []interface{}
	new     New
}

func NewPool(new New) *Pool {
	return &Pool{
		objects: make([]any, 0),
		new:     new,
	}
}

func (p *Pool) Get() any {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.objects) > 0 {
		// 队列的方式拿数据
		obj := p.objects[0]
		p.objects = p.objects[1:]
		return obj
	} else {
		return p.new()
	}
}

func (p *Pool) Put(obj interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// 队列的方式放回数据
	p.objects = append(p.objects, obj)
}

func pPool() {
	pool := NewPool(func() any {
		fmt.Println("new")
		return 1
	})
	// 获取1个对象
	x := pool.Get() // new
	fmt.Println(x)
	// 放回对象池
	pool.Put(x)
	// 放回后再次调用
	// 对象池中有对象，不用再创建
	x = pool.Get() // no "new"
	fmt.Println(x)
	// 第三次取对象
	x = pool.Get() // new
	fmt.Println(x)
}
