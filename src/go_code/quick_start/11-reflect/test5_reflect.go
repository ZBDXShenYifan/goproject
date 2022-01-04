package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (this User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Tom", 18}
	DoFileAndMethod(user)
}

func DoFileAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	for i := 0; i< inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	fmt.Println("-----------")

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}