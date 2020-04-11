//+build ignore
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s bool = true
	fmt.Println(reflect.TypeOf(s))
	var b string = strconv.FormatBool(s)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(&b)

}
