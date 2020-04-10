package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(syaHello("world"))
}
func syaHello(s string) string {
	return "Hello " + s
}
