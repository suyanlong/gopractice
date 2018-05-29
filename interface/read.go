package _interface

import "fmt"

type Read interface {
	Read() string
}

type Int int

func (i Int) Read() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) Read() string {
	return string(s)
}
