package main

import (
	"fmt"
)

func main() {
	var players = make(map[string]int)
	players["cook"] = 32
	players["angle"] = 27
	fmt.Println(players)
	fmt.Println(players["cook"])
	var aaa = make([]int, 2)
	aaa[0] = 1
	aaa[1] = 2
	fmt.Println(aaa)

	delete(players, "cook")
	fmt.Println(players)
}
