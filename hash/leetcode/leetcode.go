package main

import "fmt"

// we want to add  some value in the slice of array and get the target, 

// we need to find the index of the array ... 

func twoSum(nums []int, target int) []int {

	mapping := make(map[int]int)
	index := -1

	for i, val := range nums {
		
		if mapping == nil {
			
			index++
			mapping[val] = index
		}

		for mapval, mapindex := range mapping {
			fmt.Println("hello")

			if mapval+val == target {
				fmt.Println("target is ",target)
				return []int{i, mapindex}

			}
		}
		index++
		mapping[val] = index
	}

	return nil

}

func main() {

	nums := []int{2, 6, 2, 7, 2, 7, 4}
	fmt.Println(twoSum(nums,9))
}