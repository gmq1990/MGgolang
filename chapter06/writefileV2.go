package main

import (
	"fmt"
	"os"
)

func writefileV2() {
	err := os.WriteFile("user.txt", []byte("xxxxxxxxxx"), os.ModePerm)

	fmt.Println(err)
}
