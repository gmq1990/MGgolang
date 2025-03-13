package main

import "fmt"

func gotoCtrl() {
	var yes string
	fmt.Println("有卖西瓜的吗?(Y/n)")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")

	if yes != "Y" && yes != "y" {
		goto END
	}

	fmt.Println("买一个西瓜")

END:

	result := 0
	i := 1
START:
	// 1...100
	if i > 100 {
		goto FOREND
	}
	result += i
	i++
	goto START

FOREND:
	fmt.Println(result)
BREAKEND:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 9 {
				break BREAKEND
			}
			fmt.Println(i, j, i*j)
		}
	}

}
