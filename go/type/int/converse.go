package main

import (
	"fmt"
)

func main(){
	data := []int{1,2,3,4,5}
	
	fmt.Println(data[0:int(len(data)/2)])
	fmt.Println(data[int(len(data)/2):])
	if 2.111 > 2 {
		fmt.Println("it works that compares float with int")
	}
}

