package main

import (
	"fmt"
	"time"
	"sync"
)

var money int

func main(){
	var mux sync.Mutex
		
	go func(){
		mux.Lock()
		in(500)
		time.Sleep(100*time.Millisecond)
		see()
		mux.Unlock()
	}()
	
	go func(){
		mux.Lock()
		in(200)
		mux.Unlock()
	}()
	
	time.Sleep(200*time.Millisecond)
}

func in(m int){
	money = money + m
	fmt.Printf("i in money %d\n", money)
}

func see(){
	fmt.Printf("i have money %d\n", money)
}