package main

import (
	"fmt"
)

// we want to add  some value in the slice of array and get the target,

// we need to find the index of the array ...

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return nil // If no solution is found
}

// func twoSum(nums []int, target int) []int {

// 	mapping := make(map[int]int)
// 	index := -1

// 	for i, val := range nums {

// 		if mapping == nil {

// 			index++
// 			mapping[val] = index
// 		}

// 		for mapval, mapindex := range mapping {
// 			fmt.Println("hello")

// 			if mapval+val == target {
// 				fmt.Println("target is ",target)
// 				return []int{i, mapindex}

// 			}
// 		}
// 		index++
// 		mapping[val] = index
// 	}

// 	return nil

// }

func main() {

	// nums := []int{2, 6, 2, 7, 2, 7, 4}
	// fmt.Println(twoSum(nums, 9))

	fmt.Println(RomanInterger("VIIIL"))
}

func RomanInterger(s string) int {


	var val int

	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, char := range s {

		if value, ok := roman[string(char)]; ok {
			// fmt.Printf("character: %s,value: %d/n", string(char), value)

			val += value
			// fmt.Println("value is", value)

		} else {
			fmt.Printf("charcte: %s not found in the Roman map/n", string(char))
		}
	}

	return val

}
