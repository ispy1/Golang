//+build ignore
//reflect  反射  [riˈflekt]
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string = "ccc"
	var i int = 5
	var b bool
	f := "aaa"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(f)
}