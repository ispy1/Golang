package main

import "fmt"

func main() {

	function1()
	callback(2, Add)
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}
func Add(a, b int) {
	fmt.Println(a + b)

}
func callback(y int, f func(int, int)) {
	f(y, 2)
}
