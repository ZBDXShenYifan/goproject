package main
//四种变量声明方式
import (
	"fmt"
)
//声明一个全局变量
var gA int = 100
var gB = 200

func main() {
	//方法一
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s type of bb = %T\n", bb, bb)

	//方法三
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四 常用的方法）省去var关键字，不支持声明全局变量
	e := 100
	fmt.Println("e = ", c)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", c)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	//单行声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ",yy = ", yy)
	var kk, ll = 100, "abcd"
	fmt.Println("kk = ", kk, ",ll = ", ll)

	//多行多变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}