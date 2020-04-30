package main

import "fmt"

func main() {
	var m1 map[int]string //声明map 无空间 不能存储key
	if m1 == nil {
		fmt.Println(true)
	}

	m2 := map[int]string{}
	fmt.Println(len(m2))

	m2[4] = "red"
	fmt.Println(m2)

	m3 := make(map[int]string)
	fmt.Println(len(m3))
	m3[100] = "blue"
	fmt.Println(m3)
	fmt.Println(len(m3))

	m4 := make(map[int]string, 5) //5 长度
	fmt.Println(len(m4))
	m4[1000] = "blue"
	fmt.Println(m4)
	fmt.Println(len(m4))

	var m6 map[int]string = map[int]string{1: "2", 250: "20"}
	fmt.Println(m6)
	fmt.Println(len(m6))

	m7 := map[int]string{1: "2", 250: "20"}
	fmt.Println(m7)
	fmt.Println(len(m7))

	m8 := make(map[int]string)
	m8[100] = "aa"
	m8[30] = "bb"
	fmt.Println(m8)
	fmt.Println(len(m8))

	for k, v := range m8 {
		fmt.Println(k, v)
	}
	if v, has := m8[1]; has {
		fmt.Println(ture)
	} else {
		fmt.Println("false")
	}
}
