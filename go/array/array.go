package main

import (
	"fmt"
)

func main(){
	l := make([]int,10,10)
	fmt.Println(len(l))
	fmt.Println(cap(l))
	
	for i:=0; i<cap(l)+1; i++{
		fmt.Printf("i is %d\n", i)
		if i >= cap(l){
			l = append(l,i)
			fmt.Printf("the cap is %d\n", cap(l))
			fmt.Printf("the lenth is %d\n", len(l))
		}else{
			l[i]=i
		}
	}
	fmt.Println(l)
}

