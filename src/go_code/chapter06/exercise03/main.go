package main

import "fmt"

/*
猴子第一天吃了其中的一半
100
100 - 50

*/
func getPeach(day int) int {
	if day > 10 || day < 1 {
		fmt.Println("输入的天数有误")
		return -1
	}
	if day == 10 {
		return 1
	}else {
		return (getPeach(day+1) + 1) * 2
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(getPeach(i))
	}
}