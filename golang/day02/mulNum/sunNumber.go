package main
import (
	"fmt"
)
func SumNubers(numbers... int) int {
	total :=0
	for _, number :=range numbers{
		total += number
	}
	return total
	
}
func main()  {
	result :=SumNubers(1,2,3,45,5)
	fmt.Println(result)
	
}