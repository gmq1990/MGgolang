package main

import "os"

func file() {
	os.Rename("user.log", "user.v2.log")
	os.Remove("user.txt")
}
