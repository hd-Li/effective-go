package main

import (
	"fmt"
	"math/rand"
)

func main(){
	
	data := make([]int32, 1000)
	for i := 0;i < 1000;i++{
		num := rand.Int31n(2000)-1000
		data[i] = num
	}
	fmt.Println(data)
	data = mergeSort(data)
	fmt.Println(data)
}

//sorted_x=sort(sorted_x1, sortedx2), sorted_x1=sort(sortd_x11, sorted_x12)
func mergeSort(data []int32) []int32{
	r_data := data[0:int32(len(data)/2)]
	p_data := data[int32(len(data)/2):]
	
	if len(p_data) == 2{
		if p_data[0] > p_data[1]{
			tmp := p_data[0]
			p_data[0] = p_data[1]
			p_data[1] = tmp
		}
		
		if len(r_data) == 2{
			if r_data[0] > r_data[1]{
				tmp := r_data[0]
			    r_data[0] = r_data[1]
			    r_data[1] = tmp
			}
		}
	}else {
		r_data = mergeSort(r_data)
		p_data = mergeSort(p_data)
	}
	
	return merge(r_data, p_data)
	
}
