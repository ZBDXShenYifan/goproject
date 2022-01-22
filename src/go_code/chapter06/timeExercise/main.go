package main

import (
	"fmt"
	"time"
	"strconv"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	//test03()执行时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Println(end - start)
}