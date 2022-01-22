package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 1
	fmt.Println("i=", i)

	var j int8 = -127
	fmt.Println("j=", j)
	fmt.Printf("j的类型是 %T\n", j)
	fmt.Printf("%d\n", unsafe.Sizeof(j))
}