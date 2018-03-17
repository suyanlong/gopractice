package _interface

import "fmt"


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
