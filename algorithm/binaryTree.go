package main

import (
	"fmt"
)

var initRoot bool = false
type node struct{
	data int
	left *node
	right *node
}

func main(){
	data := [5]int{3,1,4,-2,10}
	fmt.Println(data)
	var root *node = &node{}
	for _, num := range data{
		//fmt.Printf("num is %d\n", num)
		val := &node{data:num}
		addNode(root, val)
	}
    
    binaryTreePreTraversing(root)
    fmt.Print("\n")
    binaryTreeInTraversing(root)
    
}

func addNode(root *node, val *node){
	if initRoot == false {
		root.data = val.data
		initRoot = false
	}
	
	current := root
	for {
		if val.data <= current.data {
			if current.left == nil {
				current.left = val
				return
			}
			current = current.left
			continue
		}else if val.data > current.data {
			if current.right == nil {
				current.right = val
				return
			}
			current = current.right
			continue
		}
	}
	
	return
}

func binaryTreePreTraversing(node *node) {
	fmt.Printf("%d  ", node.data)
	
	if node.left != nil {
		binaryTreePreTraversing(node.left)
	}
	
	if node.right != nil {
		binaryTreePreTraversing(node.right)
	}
	
}

func binaryTreeInTraversing(node *node) {
	if node.left != nil {
		binaryTreeInTraversing(node.left)
	}
	
	fmt.Printf("%d  ", node.data)
	
	if node.right != nil {
		binaryTreeInTraversing(node.right)
	}
	
}