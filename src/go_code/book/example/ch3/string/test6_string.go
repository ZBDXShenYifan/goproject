package main

import "fmt"
import "strconv"


func main() {
	
	/*
	将整数转化为字符串，一种是使用fmt.Sprintf,
	*/
	
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	
	
	fmt.Println(strconv.FormatInt(int64(x), 2))

	s := fmt.Sprintf("x=%b", x)

	fmt.Println(s)

	m, _ := strconv.Atoi("123")
	n, _ := strconv.ParseInt("123", 10, 64) 
	fmt.Println(m,n)
}