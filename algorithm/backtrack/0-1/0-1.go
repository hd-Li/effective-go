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

func ByBackTrack(preStones []int, preWeight int, preRecord []int){
	if len(preStones) <= 0 || preWeight <= 0{
		return
	}
	in := true
	
	if in{
		s := preStones[0]
		weight := preWeight - s
		if weight > 0{
			preRecord = append(preRecord, s)
			if len(preStones) == 1{
				fmt.Println(preRecord)
			}else{
				stones := preStones[1:]
				var record []int
				copy(record, preRecord)
				ByBackTrack(stones, weight, record)
			}
		}else{
			fmt.Println(preRecord)
		}
	}
	
	in = false
	
	if !in{
		if len(preStones) > 1{
			stones := preStones[1:]
			ByBackTrack(stones, preWeight, preRecord)
		}
	}
}
