package main

import (
	"fmt"
)

func main(){
	a := sliceMake(5)
	fmt.Println(len(a))
}

func sliceMake(n int) []int{
	a := make([]int, n)
	return a
}