package main

import "fmt"

func sortList(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := sortList(arr[:mid])
	right := sortList(arr[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))
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

// type ListNode struct {
// 	Val  int
// 	Length int
// 	Next *ListNode
// }

// func (L1 *ListNode) InsertVal(val int) {

// 	addNode := &ListNode{Val: val}

// 	if L1 == nil {

// 		L1 = addNode

// 	} else {

// 		current := L1

// 		for current.Next != nil {

// 			current = current.Next

// 		}

// 		current.Next = addNode
// 		L1.Length ++
// 	}
// }

// func (l1 *ListNode) PrintNode() {

// 	if l1 == nil {

// 		fmt.Println("there is no node more ")
// 		return

// 	}

// 	current := l1

// 	for current != nil{

// 		if current.Val != 0{
// 		fmt.Println("node data is ",current.Val)
// 		}
// 		current = current.Next
// 	}
// }

// func SortList(head *ListNode) *ListNode {

// 	if head.Length < 2 {
// 		return head
// 	}

// 	mid := head.Length/2

// 	left := len()

// }

func main() {

	arr := []int{2, 3, 10, 3, 5, 20, 6}
	fmt.Println(sortList(arr))
	

	// l1 := &ListNode{}

	// l1.InsertVal(40)
	// l1.InsertVal(20)
	// l1.InsertVal(10)
	// l1.InsertVal(30)

	// l1.PrintNode()

}
