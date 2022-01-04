package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"` 
	Sex string	`info:"sex"`
}

func findTag(str interface{}) {
	//通过 reflect.Elem() 方法获取这个指针指向的元素类型。
	//这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagstring)
		fmt.Println("doc: ", tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}