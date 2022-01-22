package main

import (
	"fmt"
)

type myFunType func(int, int) int 

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar myFunType, num1 int, num2 int) int{
	return funvar(num1, num2)
}

func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return sum, sub
}

func sum(n1 int, args... int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("%T\n", a)

	res := a(10, 40)
	fmt.Println(res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println(res2)

	res3, res4 := cal(1, 2)
	fmt.Println(res3, res4)

	res5 := sum(10, 11, 12, 13)
	fmt.Println(res5)
}