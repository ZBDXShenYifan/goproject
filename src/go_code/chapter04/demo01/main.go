package main
import (
	"fmt"
)

func main() {
	// /与%
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)

	//如果我们希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// %运算
	// 看一个公式,如果a % b = a - a / b * b        
	fmt.Println(10%3)
	fmt.Println(-10%3) // -10 - -10/3*3
	fmt.Println(10%-3) // 10 - 10/-3*-3
	fmt.Println(-10%-3)

	//++和--的使用
	var i int = 10
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
}