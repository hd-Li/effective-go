package main

import (
	"fmt"
	"time"
)

func main(){
	stop := make(chan struct{})
	
	go func(chan struct{}){
		time.Sleep(time.Second)
		fmt.Println("after sleep")
		close(stop)
	}(stop)
	
	result := <-stop
	fmt.Printf("%T\n", result)
	fmt.Println("stop the channel")
}

