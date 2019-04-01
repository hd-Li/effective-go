package main

import (
	"fmt"
)

func main(){
	fmt.Println(*foo())
}

func foo() *string {
	s := "hello world"
	return &s
}