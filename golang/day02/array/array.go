package main

import (
	"fmt"
)

func getPointer() (poniter *int) {
	a := 111
	return &a
}
func main() {
	//数组不可变
	var cheesee [2]string
	cheesee[0] = "aaa"
	cheesee[1] = "bbb"
	fmt.Println(cheesee)
	fmt.Println(cheesee[0])
	fmt.Println(cheesee[1])
	//切片 可变长度的数组
	var b = make([]string, 2)
	b[0] = "1"
	b[1] = "2"
	fmt.Println(b)
	b = append(b, "ccc", ":")
	fmt.Println(b)
	b = append(b[:1], b[1+1:]...)
	fmt.Println(len(b))
	fmt.Println(b)
	p := *getPointer()
	fmt.Println(p)
	const p1 = "aaa"
	p2 := []int{2, 3, 4}
	p2.append(p2, 1, 2, 3)
	fmt.Println(p2)

}
