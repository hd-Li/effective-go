package main

import (
	"fmt"
)

func main(){
	data1 := []int{1,3,12,33,78,91}
	data2 := []int{2,4,19,33,51,67,88,90,100}
	
	merge(data1, data2)
}

func merge(r_data, p_data []int) []int{
	r := len(r_data)
	p := len(p_data)
	l := r+p
	data := make([]int, l)
	
	i := 0
	j := 0
	for {
		fmt.Printf("i is %d\n", i)
		fmt.Printf("j is %d\n", j)
		
		if (i > r-1) || (j > p-1){
			break
		}
		
		if r_data[i] <= p_data[j]{
			data[i+j] = r_data[i]
			i++
		}else{
			data[i+j] = p_data[j]
			j++
		}
		 
	}
	
	if i > (r-1){
		for{
			if j > p-1 {
				break
			}else{
				data[i+j] = p_data[j]
				j++
			}
		}
	}else if j > (p-1){
		for{
			if i > r-1{
				break
			}else{
				data[i+j] = r_data[i]
				i++
			}
		}
	}
	
	fmt.Println(data)
	return data
}
