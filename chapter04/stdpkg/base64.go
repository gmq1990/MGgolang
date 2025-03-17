package main

import (
	"encoding/base64"
	"fmt"
)

func bBase64() {
	x := base64.StdEncoding.EncodeToString([]byte("i am tom"))
	fmt.Println(x)

	bytes, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(bytes), err)

	// base64.RawStdEncoding
	// base64.RawURLEncoding
	// base64.URLEncoding
}
