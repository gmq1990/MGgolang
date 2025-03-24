package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type User2 struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
}

func gobdecode() {
	// 文件中的内容存入内存的对象
	users := map[int]User2{
		1: User2{1, "张三", time.Now(), "123123123", "北京"},
		2: User2{2, "李四", time.Now(), "123123124", "上海"},
		3: User2{3, "王五", time.Now(), "123123125", "深圳"},
	}

	file, err := os.Open("user.gob")
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)

		fmt.Println(users)

		fmt.Printf("%#v\n", users)
	}

}
