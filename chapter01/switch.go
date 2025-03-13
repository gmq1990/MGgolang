package main

import "fmt"

func switchCtrl() {
	var yes string
	fmt.Println("有卖西瓜的吗?(Y/n)")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")

	// yes == "Y" "y"
	switch yes {
	case "y", "Y":
		fmt.Println("买一个西瓜")
	}

	fmt.Println("老公的想法：")
	switch yes {
	case "y", "Y":
		fmt.Println("买一个包子")
	default:
		fmt.Println("买十个包子")
	}

	var score int
	fmt.Println("请输入成绩：")
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
