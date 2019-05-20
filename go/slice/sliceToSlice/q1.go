package main

import (
	"fmt"
)

func main(){
	a := []int{1, 2, 3, 4, 5}
	b := a[1:]
	fmt.Println(a)
	fmt.Println(b)
	b[2] = 10
	fmt.Println(a)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	
	c := make([]int, 3)
	copy(c, a)
	fmt.Println(c)
}

