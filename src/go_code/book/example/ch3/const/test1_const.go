package main

import (
	"fmt"
	"time"
)

const pi = 3.14159
const noDelay time.Duration = 0
const timeout = 5 * time.Minute

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay) //%T输出是什么类型的数据
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println(a, b, c, d)
}

