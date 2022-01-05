package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")
		
		fmt.Println("goroutine 正在运行....")

		c <- 666 //将666发送给c
	}()

	num := <-c

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}