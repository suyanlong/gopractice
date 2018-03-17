package _interface

import "testing"

func TestInterface(t *testing.T) {
	parent := Parent{name: "suyanlong"}
	parent.Hello()
	var child Name
	child = &Child{parent}
	child.Hello()
}
