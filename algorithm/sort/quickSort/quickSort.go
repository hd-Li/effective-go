package main

import (
	"fmt"
	"math/rand"
)

func main(){
	data := make([]int, 100)
	for i := 0;i < 100;i++{
		data[i] = rand.Intn(300) - 100
	}
	fmt.Println(data)
	
	data = quickSort(data)
	
	fmt.Println(data)
}

func quickSort(data []int) []int{
	if len(data) == 1 || len(data) == 0{
		return data
	}
	mid := data[len(data)/2]
	r_data := []int{}
	p_data := []int{}
	mid_data := []int{}
	
	for i := 0;i < len(data);i++{
		if data[i] < mid{
			r_data = append(r_data, data[i])
		}else if data[i] == mid{
			mid_data = append(mid_data, data[i])
		}else if data[i] > mid{
			p_data = append(p_data, data[i])
		}
	}
	
	k := 0
	for _, num := range(quickSort(r_data)){
		data[k] = num
		k++
	}
	for _, num := range(mid_data){
		data[k] = num
		k++
	}
	for _, num := range(quickSort(p_data)){
		data[k] = num
		k++
	}
	
	return data
}

