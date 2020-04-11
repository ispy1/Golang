//+build ignore
package main

import (
	"fmt"
	"reflect"
)

//变量
/*
1使用关键是var声明一个变量
2变量名
3类型
4= 赋值运算

快捷生成声明
eg: var s,j string ="s1" ,"b1"
eg: var s string

*/
//变量

func main() {
	var s1, t1 string = "a", "b"
	var b bool = true    //布尔
	var a int = -3       //整数
	var c float64 = 3.14 //浮点数  float32 32位浮点数
	var w string = "ccc" //字符串
	var j [5]string      // 字符数组
	var y [2]int         //int数组
	var (
		s2 string = "aaa"
		t2 int    = 1
	)
	var (
		c1 int
		c2 bool
		c3 float64
		c4 string = "2"
	)
	y[1] = 1
	y[0] = 0
	j[1] = "bvbb"
	j[2] = "ccc"
	j[3] = "111"
	j[4] = "bbbbb"
	j[0] = "'wo shi'"
	fmt.Println(y)
	fmt.Println(j[0])
	fmt.Println(j)
	fmt.Println("s")
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(addition(c, 2))
	fmt.Println(w)
	fmt.Println("hello world")
	fmt.Println(syaHello("world"))
	fmt.Println(addition(2, 1))
	fmt.Println(s1, t1)
	fmt.Println(s2, t2)
	fmt.Println(c1, c2, c3, c4)
	if c4 == "" {
		fmt.Println("空")
	} else {
		fmt.Println(c4)
	}
	u := 1
	p := true
	fmt.Println(u)
	fmt.Println(reflect.TypeOf(u))
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))

}

func syaHello(s string) string {

	return "Hello " + s
}

func addition(x float64, y float64) float64 {
	return x + y
}
