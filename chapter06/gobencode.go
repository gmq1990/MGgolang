package main

import (
	"encoding/gob"
	"os"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
}

func gobencode() {
	// 内存中的内容持久化至文件
	users := map[int]User{
		1: User{1, "张三", time.Now(), "123123123", "北京"},
		2: User{2, "李四", time.Now(), "123123124", "上海"},
		3: User{3, "王五", time.Now(), "123123125", "深圳"},
	}
	file, err := os.Create("user.gob")
	if err == nil {
		defer file.Close()

		encoder := gob.NewEncoder(file)
		encoder.Encode(users)
	}
}
