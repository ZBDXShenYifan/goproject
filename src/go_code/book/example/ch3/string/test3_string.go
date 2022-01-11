package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s)) //13
	fmt.Println(utf8.RuneCountInString(s)) //9

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println("--------------------------")
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// n := 0
	// for _, _ = range s {
	// 	n++
	// }

	n := 0
	for range s {
		n++
	}
	fmt.Println(n)

	fmt.Println(utf8.RuneCountInString(s))
}
