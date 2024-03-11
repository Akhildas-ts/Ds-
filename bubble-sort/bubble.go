package bubblesort

import "fmt"

// write a code with bubble sort algoritham ..

//  arr = []int{2,35,6,2,4,6,2}
func BubbleSort(arr []int) []int {

	var temp int

	for i := len(arr); i > 0; i-- {

		for j := 1; j < i; j++ {

			if arr[j-1] > arr[j] {

				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}

	return arr

}

// get a value into the n
// get the number before the n values
// and check if the number are divisible by 3 and 5

func CheckItDivisible(n int) {

	for i := 1; i < n; i++ {

		if i % 3 == 0{
			fmt.Println("divisible by three is:",i)
		}

		if i  %5 == 0{
			fmt.Println("divisible by five is:",i)
		}

		
	}


}



