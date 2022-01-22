package main
import (
	"fmt"
)

//编写一个函数 test
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("n1=", n1)
}

//一个函数 getSum
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	//调用test
	test(n1)

	sum := getSum(10, 20)
	fmt.Println("main sum =", sum)

	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)

	res3, _ := getSumAndSub(4, 5)
	fmt.Printf("res3=%v\n", res3)
}