package main

import (
	"fmt"
)

/*switch

 */

func main() {
	i := 1
	if i >= 3 {
		fmt.Println("i is 3")
	} else if i <= 2 {
		fmt.Println("i is 2")
	} else {
		fmt.Println(i)
	}
	b := "3"
	switch b {

	case "b":
		fmt.Println("a")
	case "3":
		fmt.Println("b")
	default:
		fmt.Println("i don't know")
	}
	d := 1
	for d < 100 {
		d++
		fmt.Println(d)

	}
	for g := 0; g <= 10; g++ {
		fmt.Println(g)
	}
	num := []int{1, 2, 3, 4, 5}
	for e, n := range num {
		fmt.Println(e, n)
	}
}
