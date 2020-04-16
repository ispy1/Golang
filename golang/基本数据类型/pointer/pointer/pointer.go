package main
import(
	"fmt"
)
func main(){
	var i1=5
	fmt.Printf("An integer:%d,it's location in memory:%p\n",i1,&i1)
	if i1==*(&i1){
		fmt.Printf("true\n%")
	}
	var intp *int
	intp=&i1
	if *intp==i1{
		fmt.Println("true")
	}
}