package main

import (
	"fmt"
)

func main(){
	weight := 20
	stones := []int{1, 2, 4, 8, 2, 9, 3, 5}
	var record []int
	
	ByBackTrack(stones, weight, record)
	
}

func ByBackTrack(stones []int, weight int, preRecord []int){
	if len(stones) <= 0 || weight <= 0{
		return
	}
	//fmt.Printf("the preRecord is %v\n", preRecord)
	record0 := make([]int, len(preRecord))
	record1 := make([]int, len(preRecord))
    copy(record0, preRecord)
    copy(record1, preRecord)
	ByBackTrack(stones[1:], weight, record0)
	
	weight = weight - stones[0]
	if weight > 0{
		record1 = append(record1, stones[0])
		ByBackTrack(stones[1:], weight, record1)
	}else{
		if weight == 0{
			record1 = append(record1, stones[0])
		}
		fmt.Println(record1)
	}
}
