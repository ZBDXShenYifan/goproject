package main
import (
	"fmt"
)

func main() {
	num := new(int)
	fmt.Println(num)
	fmt.Printf("%T\n", num)
	fmt.Println(*num)
}