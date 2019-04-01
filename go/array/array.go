package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main(){
	var a [5]int
	a[1] = 1
	b := 1
	c := unsafe.Sizeof(b)
	
    fmt.Println("the type of this value is: ", reflect.TypeOf(c))
	
	fmt.Println("a is :", a)
	fmt.Println("the length of a is :", len(a))
	fmt.Println("the capacity of a is :", cap(a))
}

