package main

import (
	"fmt"
	"go-study/MGgolang/chapter05/visibility/users"
)

func main() {
	var u users.User
	// var a users.address // 未导出，包外不可见

	fmt.Printf("%#v\n", u)
	// fmt.Printf("%#v\n", a) // 未导出，包外不可见

	fmt.Println(u.ID)
	fmt.Println(u.Name)
	// fmt.Println(u.birthday) // 未导出，包外不可见
	// fmt.Println(u.addr) // 未导出，包外不可见
	fmt.Printf("%#v\n", u.Region) // address不可见，但address的Region可见
}
