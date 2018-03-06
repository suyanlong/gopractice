package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

func myFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args)
	return nil
}

func init() {
	goworker.Register("MyClass", myFunc)
}

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}



func main() {

	parent := Parent{name: "suyanlong"}
	parent.Hello()

	//child := NewChild()
	//child.Hello()

	//child2 := NewChild2()
	//child2.Hello()
}

type Namer interface {
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
	return "last name"
}

func (this *Parent) Hello() {
	fmt.Println("hello", this.GetFirstName(), this.GetLastName())
}

type Child struct {
	Namer
}

func NewChild() *Child {
	return &Child{&Parent{name: "suyanlong"}}
}

func (this *Child) GetFirstName() string {
	return "child"
}

func (this *Child) GetLastName() string {
	return "last name"
}