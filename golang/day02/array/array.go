package main

import (
	"fmt"
)

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

}
