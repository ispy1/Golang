//+build ignore
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func ShowMemAddr(x *string) {
	fmt.Println(x)
	return
}

// & 指针
func main() {
	var s bool = true
	fmt.Println(reflect.TypeOf(s))
	var b string = strconv.FormatBool(s)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(&b)
	var t string = "aaaa"
	ShowMemAddr(&t)
	fmt.Println(*&t)

}
