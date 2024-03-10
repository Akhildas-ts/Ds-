package bubblesort


 

// write a code with bubble sort algoritham .. 

//  arr = []int{2,35,6,2,4,6,2}
func BubbleSort(arr[]int)[]int{

	
	var temp int


	for i:=len(arr); i>0;i--{


		for j:=1;j<i;j++{

			if arr[j-1]>arr[j]{

				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}

	return arr
	


}