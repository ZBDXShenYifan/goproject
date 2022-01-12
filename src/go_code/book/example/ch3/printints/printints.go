package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		fmt.Println(i, v)
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']') //追加ASCII字符，若想追加任意文字符号UTF-8编码，最好用BufferRune
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}