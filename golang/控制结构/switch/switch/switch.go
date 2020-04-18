package main

import (
	"fmt"
)

func main() {
	var num int = 100
	switch num {
	case 98, 99:
		fmt.Println("98")
	case 100:
		fmt.Println("100")
	}

}
