package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Tick(time.Duration(1000000000))
	for {
		_ = <-t
		fmt.Println("123123")
	}
}
