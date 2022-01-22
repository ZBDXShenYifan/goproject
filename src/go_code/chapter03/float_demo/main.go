package main
import (
	"fmt"
)

func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)
	var num1 float32 = -0.00089
	var num2 float32 = -7809656.09

	fmt.Println("num1=", num1, "num2=", num2)

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	num8 := 5.1234e2 //5.1234 * 10^2
	num9 := 5.1234E2
	num10 := 5.1234e-2
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)

	

}