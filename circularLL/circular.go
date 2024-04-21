package main

import "fmt"

type circularLinked struct {
	head *node
}

type node struct {
	data int
	next *node
}

func (cl *circularLinked) insert(x int) {

	addNode := &node{data: x}

	if cl.head == nil {

		addNode.next = addNode
		cl.head = addNode

	} else {

		addNode.next = cl.head.next
		cl.head.next = addNode

	}
}

func (cl circularLinked) display() {

	if cl.head == nil {
		fmt.Println("Circular Linked List is empty")
		return
	}

	current := cl.head
	for {
		fmt.Printf("%d -> ", current.data)
		current = current.next
		if current == cl.head {
			break
		}
	}
	fmt.Println()
}

func main() {

	// c1 := &circularLinked{}

	// c1.insert(10)
	// c1.insert(20)
	// c1.insert(30)
	// c1.insert(40)
	// c1.insert(50)

	// c1.display()
   
	arr  := []int{2,9,3,5,13,5,2}
	sort(arr,0,len(arr)-1)

	fmt.Println(arr)
}

func sort(arr []int, low, high int) {

	if low < high {

		p1 := partion(arr, low, high)

		sort(arr, low, p1-1)
	 sort(arr, p1+1, high)

	}

}

func partion(arr []int, low, high int) int {

	piviot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {

		if arr[j] < piviot {
			i++
			arr[i], arr[j] = arr[j], arr[i]

		} else {

			arr[i+1], arr[high] = arr[high], arr[i+1]
			i++

		}

		
	}
	return i + 1

}
