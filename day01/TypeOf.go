//+build ignore
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string = "ccc"
	var i int = 5
	var b bool
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(b))
}
