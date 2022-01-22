package main
import (
	"fmt"
	"strconv"
	"unsafe"
)


//string 转基本数据类型
func main() {
	var str string = "true"
	var b bool
	//
	b, _ = strconv.ParseBool(str)
	fmt.Println(b)
	var str2 string = "123"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 8)
	fmt.Println(n1)

	var str3 string = "123.456"
	var n3 float64
	var n4 int
	n3, _ = strconv.ParseFloat(str3, 64)
	fmt.Println(n3)
	fmt.Printf("%T, %d", n4, unsafe.Sizeof(n3))
}