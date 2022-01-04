package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {

	//b: pair<type:*Book, value:book{}地址>
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	var ok bool
	w, ok = r.(Writer)
	w.WriteBook()
	fmt.Println(ok)
	fmt.Printf("%T\n", w)
	
}