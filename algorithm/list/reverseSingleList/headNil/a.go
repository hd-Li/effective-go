package main

import (
	"fmt"
)

type listNode struct{
	Val int
	Next *listNode
}

func main(){
	var list [5]*listNode
	
	for i:=0; i<5; i++{
		list[i] = &listNode{Val:i}
		fmt.Println(list[i].Val)
		if i == 0 {
			continue
		}
		
		list[i-1].Next = list[i]
		if i == 4{
			list[i].Next = nil
		}
	}
	
	head := &listNode{Next:list[0]}
	head = reverseSingleList(head)
    if head == nil{
    	return
    }
    
    current := head.Next  
    for{
    	fmt.Println(current)
    	if current.Next == nil{
    		break
    	}
    	current = current.Next
    }
	
}

func reverseSingleList(head *listNode) *listNode{
	var current, next, pre *listNode
	if head == nil{
		return nil
	}
	
	current = head.Next
	for{
		if current.Next == nil{
			current.Next = pre
			break
		}
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}
	
	head.Next = current
	
	return head
}
