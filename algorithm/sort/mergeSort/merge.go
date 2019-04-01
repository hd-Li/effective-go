package main

import(
)

func merge(r_data, p_data []int32) []int32{
	r := len(r_data)
	p := len(p_data)
	l := r+p
	data := make([]int32, l)
	
	i := 0
	j := 0
	for {
		//mt.Printf("i is %d\n", i)
		//fmt.Printf("j is %d\n", j)
		
		if (i > r-1) || (j > p-1){
			break
		}
		
		if r_data[i] < p_data[j]{
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
	
	return data
}

