package main

import (
	"fmt"
	"bytes"
)

func main() {
	s := "123456789123456789"
	fmt.Println(comma(s))
}

func comma(s string) string {
	//var buf bytes.Buffer
	b := []byte(s)
	var buf bytes.Buffer

	count := 0;
	for i := len(b)-1; i >=0; i-- {
		if (count % 3 == 0 && count > 0){
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", b[count])
		count++
	}
	
	return buf.String()

}