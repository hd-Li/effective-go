package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

func main(){
	var mux sync.Mutex
	ctx, cancel := context.WithCancel(context.Background())
	
	
	var a int
	
	go func(){
		mux.Lock()
		a = 3
		time.Sleep(3*time.Second)
		fmt.Println("after sleep")
		mux.Unlock()
	}()
	
	go func(){
		time.Sleep(100*time.Millisecond)
		mux.Lock()
		fmt.Println(a)
		mux.Unlock()
		cancel()
	}()
	
	go func(context.Context){
		time.Sleep(5*time.Second)
		fmt.Println("waiting the goroutine ends")
		<-ctx.Done()
	}(ctx)
	
	time.Sleep(6*time.Second)
}
