package main

import "fmt"

func printArray(myArray []int) {
	//引用传递
	for _, value := range myArray {
		fmt.Println("value=", value)
	}
}
func main() {
	myArray := []int{1,2,3,4}

	fmt.Printf("myArray type is %T\n", myArray)

	printArray(myArray)
}