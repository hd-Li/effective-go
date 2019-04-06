package main

import (
	"fmt"
)

func main(){
	n := 897
	fmt.Printf("%d is %b\n", n, n)
	count := countSetBitsBK(n)
	fmt.Println(count)
}

func countSetBits(n int) (count int){
	for {
		count += n & 1
		n >>= 1
		if n == 0{
			break
		}
	}
	
	return
}

func countSetBitsRecursive(n int)(count int){
	if n == 0 {
		return 0
	}
	count = n&1 + countSetBitsRecursive(n>>1)
	return
}

func countSetBitsBK(n int)(count int){
	for{
		n = n&(n-1)
		count++
		if n == 0{
			break
		}
	}
	
	return
}