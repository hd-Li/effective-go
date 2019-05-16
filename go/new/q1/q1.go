
package main

import (
	"fmt"
)

type node struct{
	val int
}

func main(){
	list := new([5]*node)
	list[0] = new(node)
	fmt.Println(list[0])
	list[0].val = 1
	fmt.Println(list[0])
}
