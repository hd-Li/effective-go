package main

import (
	"sync"
	"fmt"
)

func main(){
	var wait sync.WaitGroup
	
	wait.Add(2)
	
	go func(done func()){
		i := 0
		for{
			i++
			fmt.Println(i)
			if i == 10{
				break
			}
		}
		
		done()
	}(wait.Done)
	
	go func(done func()){
		fmt.Println("this is the second goroutine")
		done()
	}(wait.Done)
	
	wait.Wait()
}

