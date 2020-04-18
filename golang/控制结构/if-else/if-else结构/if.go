package main

import (
	"fmt"
	"runtime"
)

func main() {
	bool1:=true
	if !bool1 {
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
		if runtime.GOOS == "windows" {
			prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")		
		} else { //Unix-like
			prompt = fmt.Sprintf(prompt, "Ctrl+D")
		}
	}
	
