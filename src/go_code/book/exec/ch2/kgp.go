package main
import (
	"fmt"
	"os"
	"strconv"
	weightconv "go_code/book/exec/ch2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		kg := weightconv.Kilogram(t)
		p := weightconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			kg, weightconv.KToP(kg), p, weightconv.PToK(p))
	}
}