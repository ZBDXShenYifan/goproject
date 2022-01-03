package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	var slice2 []int
	slice2 = make([]int, 3) //开辟3个空间, 默认值是0

	slice2[0] = 100

	var slice3 []int = make([]int, 3)

	slice4 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	//判断一个slice是否为0
	var slice5 []int
	if slice5 == nil {
		fmt.Println("slice5 是一个空间切片")
	} else {
		fmt.Println("slice5 不是一个空间切片")
	}
}