package main

import (
	"fmt"
	"container/list"
)

type teacher struct{
	id int32
	name string
	titile string
}

type student struct{
	id int32
	name string
	grade int
}

func main(){
	data := list.New()
	han := teacher{id:1, name:"han", titile:"senior"}
	hd := student{id:1895, name:"hd", grade:3}
	if data.Len() == 0{
		e_han := data.PushFront(han)
		fmt.Println(data.Len())
		fmt.Println(e_han.Value)
		e_hd := data.InsertAfter(hd, e_han)
		fmt.Println(e_hd.Value)
		if e_hd.Next() == nil{
			fmt.Println("it is not a loop list")
		}
	}
}

