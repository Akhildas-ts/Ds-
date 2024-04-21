package main

import "fmt"


// DECENDING ORDER 
//MERGE SORT 

func main() {

	arr := []int{2, 2, 7, 2, 34, 5, 13, 3}

	
	result	:= mergesort(arr)
	fmt.Println(result)

}
func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {

	var result = make([]int, 0, len(left)+len(right))
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}

	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
