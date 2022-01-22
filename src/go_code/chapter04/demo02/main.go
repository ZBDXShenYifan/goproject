package main

import (
	"fmt"
)

func main() {
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok")
	}

	if age > 30 || age < 40 {
		fmt.Println("ok1")
	}

	if !(age > 40) {
		fmt.Println("ok2")
	}


	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}