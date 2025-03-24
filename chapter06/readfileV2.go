package main

import (
	"fmt"
	"os"
)

func readfileV2() {
	bytes, err := os.ReadFile("user.txt")

	if err == nil {
		fmt.Println(string(bytes))
	}
}
