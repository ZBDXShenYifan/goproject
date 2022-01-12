package main

import (
	"fmt"
	"bytes"
)

func main() {
	a := "3.1415826"
	b := "-3.1415926"
	c := "+3.1415926"
	fmt.Println(comma(a))
	fmt.Println(comma(b))
	fmt.Println(comma(c))
}

func comma(s string) string {
	//var buf bytes.Buffer
	b := []byte(s)
	var buf bytes.Buffer

	count := 0
	n := len(b)
	for i := n-1; i >=0; i-- {
		if (count % 3 == 0 && count > 0){
			buf.WriteString(",")
		}

		fmt.Fprintf(&buf, "%c", b[n-i-1])

		if (b[n-i-1] != '.' && b[n-i-1] != '+'  && b[n-i-1] != '-'){
			count++
		}
	}
	return buf.String()

}