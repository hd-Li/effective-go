package main

import (
	"math/rand"
	"fmt"
)

func main(){
	fmt.Println(rand.Int31n(10000))
	for{
		a := rand.Uint32()
		fmt.Println(a)
		if a == 888 {
			break
		}
	}
}
