package Linklist

import "fmt"




// REMOVE THE (last )N TH ELEMENT in single linked list,

//EXPLANATION
// find the length of the node
//then substract length and given element n
//then we got a postion  which element we need to delete ,
// then we are remove that element

func (l *Linked) RemoveNTHelement(head *Node, n int) {
	if head == nil {
		fmt.Println("there is nothing")
		return
	}

	current := head
	length := 1

	for current.Next != nil {

		length++
		current = current.Next

	}

	prev := head

	pos := length - n
	fmt.Println("post", pos)

	for i := 1; i < pos; i++ {
		// fmt.Println(prev.Data)
		prev = prev.Next
	}
	fmt.Println("perere", prev.Data)

	if prev == head {
		*head = *prev.Next
		return
	}
	prev.Next = prev.Next.Next

}




// PrintLinkedList prints the elements of the linked list
func (l *Linked) PrintLinkedList(head *Node) {
	if head == nil {
		fmt.Println("List is empty")
		return
	}

	current := head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
