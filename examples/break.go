package main

import "fmt"

func main() {
	count := 100
	// Add for label and break for loop
kk:
	for {
		fmt.Println(count)
		if count--; count == 0 {
			break kk
		}
	}
}
