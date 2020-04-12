package main

import (
	"fmt"
)

//函数作为参数
func antherFunc(f func() string) string {
	return f()
}
func back3(a int, b int) (int, int, int) {
	return a, b, a + b

}
func main() {
	// 函数赋值给变量
	fn := func() string {
		return "func called"
	}

	fmt.Println(antherFunc(fn))
	defer fmt.Println("ccc")
	defer fmt.Println("aaa")
	fmt.Println(back3(1, 2))
}
