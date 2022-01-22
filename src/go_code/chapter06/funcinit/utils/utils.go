package utils

import "fmt"

var Age int
var Name string

//Age 和 Name 全局变量，需要在main.go 使用
//要求初始化Age和Name

//init 函数完成初始化工作
func init() {
	fmt.Println("utils 包的 init()...")
	Age = 100
	Name = "Tom"
}