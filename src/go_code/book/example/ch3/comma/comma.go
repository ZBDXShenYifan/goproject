package main

import "fmt"

func main() {
	fmt.Println(comma("12345"))
}
/*
12345
12,345

*/
func comma(s string) string {
	fmt.Println(s)
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}