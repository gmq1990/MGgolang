package main

import "fmt"

func test() (err error) {
	defer func() {
		// recover必须在defer中使用
		if e := recover(); e != nil {
			// 用err抓取错误
			err = fmt.Errorf("%v", e)
		}
	}()
	// panic 使程序退出
	panic("error xxx")
	return

}

func panics() {
	// 使用recover抛出panic信息，使程序不中断
	err := test()
	fmt.Println(err)
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Printf("%T, %v\n", err, err) // string, error xxx
	// 	}
	// }()

	// fmt.Println("main start")
	// // panic 使程序退出
	// panic("error xxx")
	// fmt.Println("over")
}
