package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println(a[i])
	}
	s := []int{1, 2, 3, 4}
	s1 := s[:3]
	fmt.Println(s)
	fmt.Println(s1)
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4]) //'o', 'l', 'a'
	fmt.Println(b[:2])  //g o
	fmt.Println(b[2:])  //l a n g
	fmt.Println(b[:])
}
