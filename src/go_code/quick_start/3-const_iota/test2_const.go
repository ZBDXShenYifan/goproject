package main

import "fmt"

//const 定义枚举类型
//iota只能定义在const中
const (
	//可以在const() 添加iota，每行的iota都会累加1
	//第一行iota的默认值是0
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN

)

const (
	a, b = iota + 1, iota + 2 //a = 1, b = 2
	c, d //c = 2, d = 3
	e, f //e = 3, f = 4

	g, h = iota * 2, iota * 3 //g = 6, h = 9
	i, k //i = 8, k = 12
)
func main() {
	//const定义常量(只读, 不允许修改)
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a=", a, "b=", b)
	fmt.Println("c=", c, "d=", d)
	fmt.Println("e=", e, "f=", f)
	fmt.Println("g=", g, "h=", h)
	fmt.Println("i=", i, "k=", k)
}