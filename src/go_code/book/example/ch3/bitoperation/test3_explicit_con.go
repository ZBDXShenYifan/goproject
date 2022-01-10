package main

import "fmt"

func main() {
	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges //编译错误
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)

	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	f = 1e100
	i = int(f)
	fmt.Println(f, i)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}

