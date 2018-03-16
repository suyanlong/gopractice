package main

import "fmt"

func main() {
	parent := Parent{name: "suyanlong"}
	parent.Hello()
	var child Name
	child = &Child{parent}
	child.Hello()
	goroutine()
}

func goroutine() {
	//quit := make(chan int)
	//quit <- 2   无缓存的通道是一直阻塞的。
	fmt.Println("end")
	//<-quit
}

type Name interface {
	GetFirstName() string
	GetLastName() string
	Hello()
}

type Parent struct {
	name string
}

func (this *Parent) GetFirstName() string {
	return "parent"
}

func (this *Parent) GetLastName() string {
	return this.name
}

func (this *Parent) Hello() {
	this.qwe()
	fmt.Println("hello", this.GetFirstName(), this.GetLastName())
}

func (this *Parent) qwe() {
	fmt.Println("Parent")
}

type Child struct {
	Parent
}

func (this *Child) GetFirstName() string {
	return "child"
}

func (this *Child) GetLastName() string {
	return "123"
}

//func (this *Child) Hello() {
//	this.qwe()
//	fmt.Println("hello", this.GetFirstName(), this.GetLastName())
//}

func (this *Child) qwe() {
	fmt.Println("12312")
}
