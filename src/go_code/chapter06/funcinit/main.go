package main

import (
	"fmt"
	//引入包
	"go_code/chapter06/funcinit/utils"
)

var age = test()

func test() int{
	fmt.Println("test()")
	return 90
}



//通常可以在init函数中完成初始化工作

func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main().....age=",age)
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}