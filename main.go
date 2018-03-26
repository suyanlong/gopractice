package main

import (
	"fmt"
)

const (
	EVNET1 = iota + 1
	EVNET2
	EVNET3
)

func main() {
	fmt.Println(EVNET2)
	fmt.Println(EVNET3)

	var b = "1234你好"
	bb := []byte(b)
	fmt.Println(bb)

}
