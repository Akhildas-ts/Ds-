package Linklist

import "fmt"

// Node creation
//Node print
//Delete value

type Node struct {
	Data int
	Next *Node
}
type Linked struct {
	Head   *Node
	Lenght int
}

func (l *Linked) Prepend(n *Node) {

	second := l.Head

	l.Head = n
	l.Head.Next = second
	l.Lenght++
}

//Print the linkedList  data

func (l Linked) PrinDataOnNode() {

	toPrint := l.Head
	for l.Lenght != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.Lenght--

	}
	fmt.Printf("\n")
}

//deleting the value from the linked List

func (l *Linked) DeleteLinkedListData(value int) {

	previousToDelete := l.Head

	if l.Head == nil {
		return

	}
	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Lenght--
		return
	}

	for previousToDelete.Next.Data != value {

		if previousToDelete.Next.Next == nil {
			fmt.Println("there is no data for delete")
			return
		}
		previousToDelete = previousToDelete.Next
	}

	previousToDelete.Next = previousToDelete.Next.Next
	l.Lenght--
}
