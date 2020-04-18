package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var newS string
	fmt.Printf("the size of ints is:%d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer", orig)
		os.Exit(1)
	}
	fmt.Printf("the integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("the new string is %s\n", newS)
}
