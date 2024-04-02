package main

import "fmt"

// TIME COMPLEXITY OF QUICK SORT
// *BEST CASE O(LOG N)
// *AVERAGE CASE O(LOG N)
// WORST CASE O(N * 2)

// QUICK SORT ALOGRITHAM
// frist of all we got a array then we need to take any index of value from the array , that value called piviot
// then we need to check   less than that value and greater than that value .  so that time we got the correct postion of the value,
// then we  work resursily , frist will arange all the left side , in same concept , and after we we arrange the right side .
func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {

	arr := []int{2, 3, 5, 8, 7}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// QUICK SORTING WITH AGE OF STUDENT'S

// type Student struct {
// 	Name  string
// 	Age   int
// 	Grade float32
// }

// func StudentSort(arr []Student) {

// 	low := 0

// 	high := len(arr) - 1

// 	QuickSort(arr, low, high)

// 	fmt.Println(arr)

// }

// func QuickSort(arr []Student, low, high int) {
// 	if low < high {

// 		p1 := partition(arr, low, high)

// 		QuickSort(arr, low, p1-1)
// 		QuickSort(arr, p1+1, high)

// 	}

// }

// func partition(arr []Student, low, high int) int {

// 	piviot := arr[high]

// 	i := low - 1

// 	for j := low; j < high; j++ {

// 		if arr[j].Age < piviot.Age {
// 			i++

// 			arr[i], arr[j] = arr[j], arr[i]

// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	return i + 1
// }

// func LargestElementK(arr []int, k int) {

// 	low := 0

// 	high := len(arr) - 1 // Adjusted here

// 	quickSort(arr, low, high)

// 	p := len(arr) -k

// 	fmt.Println(arr[p])

// }
// func quickSort(arr []int, low, high int) {

// 	if low < high {
// 		p := partition(arr, low, high)

// 		quickSort(arr, low, p-1)
// 		quickSort(arr, p+1, high)
// 	}

// }

// func partition(arr []int, low, high int) int {
// 	piviot := arr[high]
// 	i := low - 1

// 	for j := low; j < high; j++ {

// 		if arr[j] < piviot {
// 			i++

// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	return i + 1
// }

//

// ****
