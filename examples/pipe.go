package main

import "io"

func main() {
	read, write := io.Pipe()
	go write.Write([]byte("123333"))
	r := make([]byte, 100)
	read.Read(r)
	println(string(r))
}
