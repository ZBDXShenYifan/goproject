package main

import "fmt"

func test() {
	//使用defer + recover来捕获和处理异常
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	test()
	fmt.Println("main...")
}