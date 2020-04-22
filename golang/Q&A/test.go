package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println(a[i])

	}
	b := 10
	if b == *&b {
		fmt.Print(1)
	}
}
