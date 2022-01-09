package main

import (
	"flag"
	"fmt"
	"strings"
)


//flag.Bool函数创造一个新的布尔标识变量。
var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator") //-s sep 使用sep替换默认参数输出时的空格换行符

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}