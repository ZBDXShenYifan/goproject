package main
import (
	"fmt"
	"os"
	"strconv"
	lenconv "go_code/book/exec/ch2/lenconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		in := lenconv.Inch(t)
		m := lenconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			in, lenconv.InToM(m), m, lenconv.MToIn(in))
	}
}