package main
import (
	"fmt"
)

func main() {
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期,%d天\n", week, day)

	//program2
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v转摄氏度是%v", huashi, sheshi)
}