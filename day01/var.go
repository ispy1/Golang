//+build ignore
package main

import (
	"fmt"
)

//变量
func main() {

	var b bool = true    //布尔
	var a int = -3       //整数
	var c float64 = 3.14 //浮点数  float32 32位浮点数
	var w string = "ccc" //字符串
	var t [5]string      // 字符数组
	var y [2]int         //int数组
	y[1] = 1
	y[0] = 0
	t[1] = "bvbb"
	t[2] = "ccc"
	t[3] = "111"
	t[4] = "bbbbb"
	t[0] = "'wo shi'"
	fmt.Println(y)
	fmt.Println(t[0])
	fmt.Println(t)
	fmt.Println("s")
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(addition(c, 2))
	fmt.Println(w)
	fmt.Println("hello world")
	fmt.Println(syaHello("world"))
	fmt.Println(addition(2, 1))

}

func syaHello(s string) string {
	return "Hello " + s
}

func addition(x float64, y float64) float64 {
	return x + y
}
