package main

import (
	"fmt"
	"sync"
	"bytes"
	"io"
	"io/ioutil"
	"time"
)

func main(){
	var mu sync.Mutex
	var buffer bytes.Buffer
	sign := make(chan struct{}, 2)
	
	
	go func(writer io.Writer){
		defer func(){
			sign <- struct{}{}
		}() 
		
		head := fmt.Sprint("i am hd\n")
	    body := fmt.Sprint("i am wd\n")
	    mu.Lock()
		writer.Write([]byte(head))
		writer.Write([]byte(body))
		mu.Unlock()
	}(&buffer)
	
	test := fmt.Sprint("test\n")
	mu.Lock()
	(&buffer).Write([]byte(test))
	time.Sleep(3*time.Second)
	(&buffer).Write([]byte("test2\n"))
	mu.Unlock()
	<-sign
	data, _ := ioutil.ReadAll(&buffer)
	fmt.Printf("%s", data)
}
