package main

import (
	"fmt"
	"time"
)

func main() {
	k := time.After(time.Second)
	kk := time.Tick(time.Second * 2)
	for {
		select {
		case <-k:
			fmt.Println(123)
		case <-kk:
			fmt.Println(12323)
		}
	}
}
