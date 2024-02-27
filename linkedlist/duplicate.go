package Linklist

import "fmt"

func RemoveDuplicates(head *Node) *Node {
	current := head
	for current != nil && current.Next != nil {
		if current.Data == current.Next.Data {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
func DisplayList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
