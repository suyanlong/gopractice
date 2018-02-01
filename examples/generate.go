package main

import "fmt"

//go:generate echo hello, line=$GOLINE
//go:generate go run generate.go
//go:generate echo file=$GOFILE, pkg=$GOPACKAGE
func main() {
	fmt.Println("hello world~")
}
