package main

import (
	"fmt"
)

type cat struct {
	Name string
	Age  int
}

//嵌套结构体
type HomeCat struct {
	cat   cat
	color string
}

func main() {
	a := cat{
		Name: "doujiang",
		Age:  6,
	}
	var b cat
	b.Name = "bbb"
	b.Age = 3
	fmt.Println(a)
	fmt.Println(a.Age, a.Name)
	fmt.Println(b)
	c := cat{
		Name:"aaa",
		Age:7,
	}
	e:=cat{
		Name:"doujiang",
		Age:8,
	}
	fmt.Println(c)
	d := HomeCat{
		color: "red",
		cat: cat{
			Name: "aaa",
			Age:  1,
		},
	}
	fmt.Println(d)
	fmt.Println(d.cat.Age)
	if a==e{
		fmt.Println("true")
	}else{
		fmt.Println("flase")
	}
	f:=&a
	f.Age=8
	fmt.Println(*&f)
	fmt.Printf("%p\n",&a)
	fmt.Println(f)
	fmt.Println(a)
	
}
