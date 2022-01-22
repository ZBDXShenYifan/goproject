package main
import (
	"fmt"
)

func main() {
	var i int32 = 100
	var n float32 = float32(i)
	fmt.Println(n)

	//测试题
	var n1 int32 = 20
	var n2 int64
	var n3 int8
	
	n2 = int64(n1) + 20 //int32赋给int64
	n3 = int8(n2) + 20 //int64赋给int8


	fmt.Println(n1, n2, n3)
}