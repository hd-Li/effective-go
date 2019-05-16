package main

import (
	"container/list"
	"fmt"
)

func main(){
	root := list.New()
	a := list.Element{Value:1}
	root.PushBack(a)
	root.PushFront(list.Element{Value:0})
	fmt.Println(root.Back())
	fmt.Println(root.Front())
	fmt.Printf("the type of value is %T", root.Front().Next().Value)
}
