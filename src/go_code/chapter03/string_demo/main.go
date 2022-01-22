package main
import (
	"fmt"
)

func main() {
	var address string = "北京长城 110 hello world!"
	fmt.Println(address)

	var str = "hello"
	var str2 = `abc\nabc`
	fmt.Println(str2, str)

	str = "hello" + " world!"
	str += "haha"
	fmt.Println(str)

	var a int
	var b float64
	var c bool
	var d string
	

	fmt.Println(a,b,c,d)
}