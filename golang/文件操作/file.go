package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("d:/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("succeed")

}
