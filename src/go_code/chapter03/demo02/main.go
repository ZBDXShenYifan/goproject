package main
import "fmt"

//演示golang中+的使用
func main() {

	var i = 1
	var j = 2
	var r = i + j
	fmt.Println(r);

	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2 // 拼接操作
	fmt.Println(res)
}