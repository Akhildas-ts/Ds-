package Linklist

import "fmt"

func (l *Linked) InsertAfter(x, data int) {

	newNode := &Node{Data: data, Next: nil}

	current := l.Head

	for current != nil {
		if current.Data == x {
			newNode.Next = current.Next
			current.Next = newNode
			l.Lenght++

		}

		current = current.Next

	}
}

func (l *Linked) InsertBefor(x, data int) {

	newNode := &Node{Data: data, Next: nil}

	if l.Head.Data == x {

		newNode.Next = l.Head

		l.Head = newNode
		l.Lenght++
	}
	prev := l.Head
	current := l.Head.Next

	for current != nil {

		if current.Data == x {
			prev.Next = newNode
			newNode.Next = current
			l.Lenght++
			return
		}
		prev = current
		current = current.Next

	}

	fmt.Printf("node  with data not found \n", x)

}



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
