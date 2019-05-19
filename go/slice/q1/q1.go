package main

import (
	"fmt"
	"container/list"
)

func main(){
	//l := [5]*list.List
	l := make([]*list.List, 0, 5)
	fmt.Println(l[0].Front())
}

func sliceMake(n int) []int{
	a := make([]int, n)
	return a
}