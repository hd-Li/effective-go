package main

import (
	"fmt"
	"container/list"
)

func main(){
	slice := make([]int,2)
	l := list.New()
	a := 1
	l.PushBack(a)
	e := l.Front()
	slice[e.Value.(int)] = 10
	fmt.Println(slice)
	fmt.Printf("%T\n", e.Value)
	if l.Front().Prev() == nil{
		fmt.Println("the next of back is nil")
	}
}

