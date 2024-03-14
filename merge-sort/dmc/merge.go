package main

import "fmt"

func main() {

	arr := []int{2, 3, 24, 1, 16, 5}

	result := MergeSort(arr)
	fmt.Println("merged array is ,",result)
}



// MERGE SORT 

// FRIST OF ALL WE WANT TO KNOW ABOUT THE CONCEPT OF THE RECURSION , 
// GET A ARRAY FIND MID VALUE 
// SPLIT TWO PART MID BEFORE AND MID AFTER ,
// THEN WE SPLIT RECURSLIVIY THEN THEN LAST WE GOT TWO ELEMENT THEN ARRANGE ACCENDING ORDER \
// THEN FUNC GO BACK RECUSILY , ITS BETTER TO DO IT ON PAPPER, 

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left []int, right []int) []int {

	var result = make([]int,0, len(left)+len(right))

	i := 0
	j := 0

	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}

	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}
