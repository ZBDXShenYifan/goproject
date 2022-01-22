package main
import (
	"fmt"
	"strconv"
)

func main() {

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	//使用第一种方式进行转换
	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)
	str = fmt.Sprintf("%f", num2)
	fmt.Println(str)
	str = fmt.Sprintf("%t", b)
	fmt.Println(str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Println(str)

	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Println(str)
}