package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["1"] = "111"
	m["2"] = "222"

	for value := range m {
		fmt.Println(value)
	}
}
