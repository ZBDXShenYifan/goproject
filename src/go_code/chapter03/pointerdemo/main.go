package main
import (
	"fmt"
)

//演示golang的数据类型
func main() {
	
	//基本数据类型在内存的布局
	var i  int = 10
	//i的地址
	fmt.Println("i的地址是=", &i)
	var ptr *int = &i;
	fmt.Printf("ptr=%v", ptr)
	fmt.Printf("ptr的地址=%v", &ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)

	var num int = 9
	fmt.Println("num address=%v", )

	var ptr2 *int
	ptr2 = &num
	*ptr2 = 10
	fmt.Printf("num=%v", num)
}