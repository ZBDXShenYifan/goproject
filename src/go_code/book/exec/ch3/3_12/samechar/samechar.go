package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(samechar("abc", "abc"))
	fmt.Println(samechar("6428", "6237"))
	fmt.Println(samechar("6428", "6428"))
}

func samechar(s1, s2 string) bool {
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}