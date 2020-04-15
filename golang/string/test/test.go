package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main(){
	str1:="  sdas 123 321 1  s"
	fmt.Printf("%d\n",len(str1))
	fmt.Printf("%d\n",utf8.RuneCountInString(str1))
	a:=strings.HasPrefix(str1,"d")
	fmt.Println(a)
	fmt.Println(strings.HasSuffix(str1,"23"))
	fmt.Println(strings.HasSuffix(str1,"223"))
	fmt.Println(strings.Contains(str1,"12"))
	fmt.Println(strings.Index(str1 ,"12"))
	fmt.Println(strings.LastIndex(str1,"21"))
	fmt.Println(strings.Replace(str1,"3","aa",1))
	fmt.Println(strings.Count(str1,"3"))
	fmt.Println(strings.Repeat(str1,5))
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.TrimSpace(str1))
	fmt.Println(strings.TrimRight(str1,"s"))
	fmt.Println(strings.TrimLeft(str1,"s"))
	
}
