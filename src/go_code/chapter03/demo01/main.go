package main
import "fmt"

// var a = "Jack"

// var (
// 	b = "Array"
// 	c = "Marry"
// 	d = "Colina"
// )

func main() {
	// var i int = 10
	// fmt.Println("i=", i)

	// //int 默认为0
	// var j int
	// fmt.Println("j=", j)

	// //根据值自动判断数据类型
	// var num = 10.11
	// fmt.Println("num=", num)

	//省略var, 注意 := 左侧不应该是已经声明过的
	//等价于 var name string 
	//      name = "tom"

	//name := "tom"
	//fmt.Println("name=", name)

	// var n1, n2, n3 int
	// fmt.Println("n1=",n1, "n2=", n2, "n3=", n3)

	// var n1, name, n3=100, "tom", 888
	// fmt.Println("n1=", n1, "name=", name, "n3", n3)

	// n4, n5, n6 := 1001, 1002, 1003
	// fmt.Println("n4=", n4, "n5=", n5, "n6=", n6)

	//fmt.Println("a=", Jack)
	//fmt.Println("b=", b, "c=", c, "d=", d)

	var i = 30
	i = 20
	i = 5
	fmt.Println("i=", i)
	//i = 1.2 //int 原因是不能更改数据

	// i := 99 不能重复定义

	//变量不能重复定义
	//Golang会有一些默认初始值， int 0, string 为空串
	

}