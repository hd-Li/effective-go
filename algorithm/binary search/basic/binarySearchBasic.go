package main

import (
	"fmt"
	"math/rand"
)

func main(){
	data := make([]int, 40, 40)
	start := 0
	for i:=0;i<len(data);i++{
		start += rand.Intn(100)
		data[i] = start
	}
	fmt.Println(data)
	dataTest := []int{9}
	got := binarySearch(dataTest, 9)
	fmt.Println(got)
}

func binarySearch(data []int, target int)(got bool){
	if len(data) == 0{
		fmt.Println("invalid data:length is 0")
		return false
	}
	start := 0
	end := len(data) - 1
	for{
		//mid := start + (end-start)/2
		mid := start + ((end - start) >> 1)
		fmt.Printf("mid is %d\n", mid)
		if data[mid] == target{
			return true
		}else if data[mid] < target {
			start = mid + 1
		}else{
			end = mid - 1
		}
		if start > end {
			break
		}
	}
	return false
}

func binarySearchRecursiveOnly(data []int, target int)(got bool){
	if len(data) == 0{
		err := fmt.Errorf("valid data: slice is null")
		fmt.Println(err.Error())
		return false
	}
	fmt.Printf("data is %v\n", data)
	start := 0
	end := len(data) -1
	mid := start + (end-start)/2
	fmt.Printf("mid is %d\n", mid)
	var got1, got2 bool
	if data[mid] == target {
		return true
	}else {
		if start <= (mid-1) {
			got1 = binarySearchRecursiveOnly(data[start:(mid-1)], target)
		}
		
		if (mid+1) <= end {
			got2 = binarySearchRecursiveOnly(data[(mid+1):(end+1)], target)
		}
		
		return (got1 || got2)
	}
	
	return false
}

func binarySearchRecursive(data []int, target int) (got bool){
	if len(data) == 0{
		err := fmt.Errorf("valid data: slice is null")
		fmt.Println(err.Error())
		return false
	}
	
	start := 0
	end := len(data) -1
	mid := start + (end-start)/2
	fmt.Printf("mid is %d\n", mid)
	if data[mid] == target {
		return true
	}else if data[mid] < target {
		start = mid + 1
	}else{
		end = mid - 1
	}
	
	if start <= end{
		data = data[start:(end+1)]
		fmt.Printf("data is %v\n", data)
		got = binarySearchRecursive(data, target)
		
		return got
	}
	
    return false	
}