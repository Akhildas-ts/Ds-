package main

import "fmt"

func main() {
	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 3, 12, 52, 3, 5, 6, 2}
	for _, val := range buildHeap {

		m.Insert(val)
		fmt.Println(m)

	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}

//Max Heap ..

// Max heap struct has a slice that holds the array

type MaxHeap struct {
	array []int
}

//Insert Add an Element to the heap

func (h *MaxHeap) Insert(key int) {

	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array) - 1)

}

// Extract the largest key and remove key, and remove it from the heap ...
func (h *MaxHeap) Extract() int {

	if len(h.array) == 0 {
		fmt.Println("there is no more elements in the  array ")
		return -1
	}
	exrtact := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	
	h.maxHeapifydown(0)

	return exrtact

}

// Extract return  the largest key, and removes it from the heap

func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {

		h.swap(parent(index), index)

		index = parent(index)
	}

}

// maxHeapifydown  will heapily top  to bottom
func (h *MaxHeap) maxHeapifydown(index int) {
	lastIndex := len(h.array) - 1

	
	l, r := left(index), right(index)

	childCompare := 0

	// loop  while index  has at least one child
	for l <= lastIndex {

		if l == lastIndex { //when the left child is only there
			childCompare = l
		} else if h.array[l] > h.array[r] { // when the left child is larger
			childCompare = l
		} else { // when right child is larger
			childCompare = r
		}

		//when the child value is greater than parent then it swap 
		if h.array[index] < h.array[childCompare] {

			 // swap the array value in these case
			h.swap(index, childCompare)

			//assign the index value 
			index = childCompare
			l, r = left(index), right(index)

		} else {
			return
		}

	}
}

// parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left child index
func left(i int) int {

	return 2*i + 1
}

// right chid index
func right(i int) int {
	return 2*i + 2
}
func (h *MaxHeap) swap(i1, i2 int) {

	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
