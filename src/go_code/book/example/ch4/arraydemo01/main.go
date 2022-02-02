package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	//定义一个数组
	var a [6]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2])
	fmt.Println(r[2])

	q1 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q1)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "a", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	//r := [...]int{99: -1}
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1)
	
	
}