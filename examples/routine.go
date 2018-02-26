package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	// CPU 占用for循环。
	//for {}
	fmt.Println("hello world!")
}
