package main

import (
	"fmt"
	"unicode/utf8"
)

func utf_8() {
	s := "现在我有冰淇淋"
	// len() 是字节的长度，不是rune的长度
	fmt.Println(len(s)) // 21
	// utf8的runecount
	fmt.Println(utf8.RuneCountInString(s)) // 7
}
