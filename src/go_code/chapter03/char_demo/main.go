package main
import (
	"fmt"
)

func main() {

	var c1 byte = 'a'
	var c2 byte = '0'

	//当直接输出byte值时，就是输出了对应的字符的ASCII码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c, c2=%c\n", c1, c2)

	var c3 int ='北'
	fmt.Printf("c3=%c\n", c3)
}