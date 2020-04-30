package main

import "fmt"

//Person 数据结构
type Person struct {
	name string
	sex  byte
	age  int
}

func main() {
	var man Person = Person{name: "bob", sex: 'm', age: 18}
	fmt.Println(man)
	m2 := Person{sex: 'f'}
	fmt.Println(m2.age)

	var man3 Person
	man3.name = "aaa"
	man3.age = 11
	fmt.Println(man3)
	if man == m2 {

	}
}
