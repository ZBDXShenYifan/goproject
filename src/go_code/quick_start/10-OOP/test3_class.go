package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()....")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()....")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	s := SuperMan{Human{"li4", "female"}, 88}

	s.Walk()
	s.Eat()
	s.Fly()
	s.Print()

	var k SuperMan
	k.name = "wang5"
	k.sex = "man"
	k.level = 3
	k.Print()
}