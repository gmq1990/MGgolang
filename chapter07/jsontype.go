package main

import (
	"encoding/json"
	"fmt"
)

const (
	Large  = iota // large
	Medium        // medium
	Small         // small
)

// 自定义类型
type Size int

func (s Size) MarshalText() ([]byte, error) {
	switch s {
	case Large:
		return []byte("large"), nil
	case Medium:
		return []byte("medium"), nil
	case Small:
		return []byte("small"), nil
	default:
		return []byte("unknown"), nil
	}
}

// 方法名首字母必须大写，否则json不能识别
func (s *Size) UnmarshalText(bytes []byte) error {
	switch string(bytes) {
	case "large":
		*s = Large
	case "medium":
		*s = Medium
	case "small":
		*s = Small
	default:
		*s = Small
	}
	return nil
}

func jsontype() {
	// 自定义json序列化
	var size Size = Medium

	bytes, _ := json.Marshal(size)
	fmt.Println(string(bytes)) // "medium"

	// 自定义json反序列化
	var size02 Size
	json.Unmarshal(bytes, &size02)
	fmt.Println(size02) // 1

	sizes := []Size{Large, Large, Small, Medium}
	bytes, _ = json.Marshal(sizes)
	fmt.Println(string(bytes)) // ["large","large","small","medium"]

	var sizes02 []Size
	json.Unmarshal(bytes, &sizes02)
	fmt.Println(sizes02) // [0 0 2 1]
}
