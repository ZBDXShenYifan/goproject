package main
import (
	"fmt"
)

func main() {

	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("Please enter the age")
	fmt.Scanln(&age)
	fmt.Println("Please enter the salar")
	fmt.Scanln(&sal)
	fmt.Println("Please enter the result of isPass")
	fmt.Scanln(&isPass)
	fmt.Println(name, age, sal, isPass)

	fmt.Println("请输入姓名，年龄，薪水，是否通过")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Println(name, age, sal, isPass)
}