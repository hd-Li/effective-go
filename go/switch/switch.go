package main

import (
	"fmt"
)

func main(){
	fmt.Println("app started")
	
	switch a := 3; a {
		case 1 :
		  fmt.Println("the number is :", a)
        case 2 :
          fmt.Println("yes")		  
	}
	
	b := 2*5
	fmt.Println("b is :", b)
}

