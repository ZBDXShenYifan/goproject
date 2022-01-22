package main

import (
	"fmt"
	"go_code/chapter06/funcdemo01/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 3.2
	var operator byte = '+'
	result := utils.Cal(n1, n2, operator)
	fmt.Println("result=", result)

	
}

