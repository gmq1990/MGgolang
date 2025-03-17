package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func mD5() {
	// 几个方法，结果都相同
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte("i am tom")))
	fmt.Println(md5Str)

	bytes := md5.Sum([]byte("i am tom"))
	fmt.Println(hex.EncodeToString(bytes[:]))

	m := md5.New()
	m.Write([]byte("i am"))
	m.Write([]byte(" tom"))
	fmt.Printf("%x\n", m.Sum(nil))

}
