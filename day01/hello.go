package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(syaHello("world"))
	fmt.Println(addition(2, 1))
}
func syaHello(s string) string {
	return "Hello " + s
}

func addition(x int, y int) int {
	return x + y
}
