package main

import (
	"fmt"
)

type node struct{
	next *node
	employee *employee
} 

type employee struct{
	id int32
	name string
	titile string
}
const size int = 100

func main(){
	head := &node{}
	l := 0
	
	hd := &employee{id:1895, name:"hd", titile:"intermediate"}
	head.next = &node{employee: hd}
	l++
	fmt.Println(head.next.employee)
}

