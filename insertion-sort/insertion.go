package insertionsort

//Insertion sorting
// we find the length of array
// frist the loop will have one step front .
// we create a key varaible for store the value of i
// then  we create a anothere loop its index was one step back of frist one

func InsertionSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // Move elements greater than key one position ahead
			j = j - 1
		}
		arr[j+1] = key // Insert key at its correct position
	}

}
