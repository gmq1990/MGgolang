package main

import (
	"bytes"
	"fmt"
)

func bBytes() {
	var bytes1 []byte = []byte{'a', 'b', 'c'}
	fmt.Printf("%T, %#v\n", bytes1, bytes1) // []uint8, []byte{0x61, 0x62, 0x63}

	// 字节切片 => 字符串
	s := string(bytes1)
	fmt.Printf("%T %v\n", s, s)

	// 字符串 => 字节切片
	bs := []byte(s)               // string abc
	fmt.Printf("%T %v\n", bs, bs) // []uint8 [97 98 99]

	fmt.Println(bytes.Compare([]byte("abc"), []byte("def")))         // -1
	fmt.Println(bytes.Index([]byte("abcdefabc"), []byte("def")))     // 3
	fmt.Println(bytes.Contains([]byte("abcdefabc"), []byte("defd"))) // false

}
