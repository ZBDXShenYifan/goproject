package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	// c := s[len(s)] 下标越界
	fmt.Println(s[0:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	s1 := "left foot"
	t := s1
	s1 += ", right foot"
	fmt.Println(s1)
	fmt.Println(t)
	//s[0] = 'L' 错误，字符串的值不可修改

	s2 := "hello, world"
	hello := s[:5]
	world := s[7:]
	//s2 hello world 共用底层字节数组
	fmt.Println(s2, hello, world)

	//原生的字符串字面量
	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
		go command [arguments]
	...`
	fmt.Println(GoUsage)
}