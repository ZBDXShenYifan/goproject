package main

import "fmt"
import "math"

func main() {
	var f float32 = 16777216 //1<<24
	fmt.Println(f == f+1)
	
	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	fmt.Println(e, Avogadro, Planck)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

}

// func compute() (value float64, ok bool){
// 	//...
// 	if faild {
// 		return 0, false
// 	}
// 	return result, true
// }