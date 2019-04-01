package main

import (
	"fmt"
)

func main(){
	a := make([]int,5)

    a = []int{1,2}
    fmt.Println(a)
    
    b := []int{}
    b = append(b,1)
    fmt.Println(b)
   // c := make([]int, 3)
    d := []int{}
    for _, num := range(d){
    	fmt.Println(num)
    }
  
}

