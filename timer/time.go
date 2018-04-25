package main

import (
	"fmt"
	"time"
)

func main() {
	k := time.After(time.Second * 10)
	kk := time.Tick(time.Second * 2)
	for {
		select {
		case <-k:
			fmt.Println(12311111111111)
		case <-time.After(time.Second * 3):
			fmt.Println(123)
		case <-kk:
			fmt.Println(12323)
		}
	}
}
