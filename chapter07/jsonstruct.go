package main

import (
	"encoding/json"
	"fmt"
)

// 1. json序列化，要求元素的首字母大写（包外可见）
type User struct {
	ID   int    `json:"-"` // 属性的标签。:的前面为标签名，:后为标签值，-表示忽略
	Name string `json:"name"`
	Sex  int    `json:"sex,int,omitempty"` // omitempty表示：0值时不显示结果
	tel  string
	Addr string
}

func jsonStruct() {
	user := User{1, "tom", 0, "123123123123", "beijing"}

	bytes, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(bytes)) //   tel 不显示

	var user02 User

	json.Unmarshal(bytes, &user02)
	fmt.Println(user02) // bytes没有tel，反序列化后user02没有tel属性
}
