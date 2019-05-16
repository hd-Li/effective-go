package main

import (
	"fmt"
)

type ListNode struct{
	Val int
	Next *ListNode
}

func main(){
	a := &ListNode{Val:5}
	b := &ListNode{Val:4}
	c := &ListNode{Val:3}
	d := &ListNode{Val:2}
	e := &ListNode{Val:1}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next =  nil
    head := a
	 
	head = reverseList(head)
	
	for{
		fmt.Println(head.Val)
		if head.Next == nil{
			break
		}
		head = head.Next
	}
	
}

func reverseList(head *ListNode) *ListNode {
    var current, pre, next *ListNode
    
    if head == nil{
        return nil
    }
    current = head
    
    for{
        if current.Next == nil {
        	current.Next = pre
            break
        }
        next = current.Next
        current.Next = pre
        pre = current
        current = next
    }

    return current
}