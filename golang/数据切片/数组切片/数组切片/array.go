package main

import (
	"fmt"
)

func main() {
	var arr1 = new([5]int)
	arr2 := *arr1
	arr2[2] = 100
	fmt.Println(arr2)
	fmt.Println(*arr1)
}
