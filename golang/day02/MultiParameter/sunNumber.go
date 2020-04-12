package main

import (
	"fmt"
)

func SumNubers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}
func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}
func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten == 5 {
		fmt.Printf("i'm full i've eaten %d \n", eaten)
		return eaten
	}
	fmt.Println("i'm still hungry")
	return feedMe(portion, eaten)
}

func main() {
	result := SumNubers(1, 2, 3, 45, 5)
	fmt.Println(result)
	fmt.Println(sayHi())
	fmt.Println(feedMe(1, 0))
	fn := func() {
		fmt.Println("called")
	}
	fn()
}
