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
	// gob文件中的内容存入内存的对象
	users := map[int]User2{}

	file, err := os.Open("user.gob")
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)

		fmt.Println(users)

		fmt.Printf("%#v\n", users)
	}

}
