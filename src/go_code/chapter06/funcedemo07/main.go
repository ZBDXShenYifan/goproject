package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += "a"
		fmt.Println("str=" + str)  //1. str="helloa" 2. str="helloaa" 3. str="helloaaa"
		return n
	}
}

func makeSuffix(suffix string) func (string) string {
	return func (name string) string {
		//如果 name 没有指定的后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	//makeSuffix
	//返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("bird.avi"))
}