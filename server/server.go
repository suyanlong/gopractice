package server

import (
	"io"
	"fmt"
)

///匿名接口，匿名结构体组合，都值得推敲
///四种组合见网址解释： https://studygolang.com/articles/6934

type Server interface {
	io.Closer
	server()
}

type server struct {
	Server //匿名接口
}

type base struct {
}

func (b *base) server() {
	fmt.Println("server")
}

func (b *base) Close() error {
	fmt.Println("close base")
	return nil
}

func service() {
	var s Server
	s = server{&base{}}
	s.Close()
}

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	fmt.Println("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func S() {
	user := User{
		Name:  "AriesDevil",
		Email: "ariesdevil@xxoo.com",
	}

	SendNotification(&user)
	//SendNotification(&user)
}
