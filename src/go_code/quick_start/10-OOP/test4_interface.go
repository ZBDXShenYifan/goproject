package main

import "fmt"
//多态
//本质是一个指针
//有一个父类(有接口)
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}


//有不同的子类(实现了父类的全部接口方法)
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())

}

func main() {
	// var animal AnimalIF
	// animal = &Cat{"Green"}
	
	// animal.Sleep()

	// animal = &Dog{"Yello"}

	// animal.Sleep()

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	//父类类型的变量（指针）指向（引用）子类的具体数据变量
	showAnimal(&cat)
	showAnimal(&dog)

}