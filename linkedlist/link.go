package Linklist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {

	second := l.head
	l.head = n
	l.head.next = second
	l.Length--
}

func (l *LinkedList) Delete(value int) {

	previousToHead := l.head

	if previousToHead.data == value {
		return
	}

	for previousToHead.next.data != value {

		if previousToHead.next.next == nil {
			fmt.Println("there is no data for delete")
			return
		}

		previousToHead = previousToHead.next

	}

	previousToHead.next = previousToHead.next.next
}

func (l LinkedList) PrinDataOnNode() {
	headTOPrint := l.head

	for l.Length != 0 {
		fmt.Printf("%d", headTOPrint.data)
		headTOPrint = headTOPrint.next
		l.Length--
	}
}

// type Node struct {
// 	Data int
// 	Next *Node
// }
// type Linked struct {
// 	Head   *Node
// 	Lenght int
// }

// func (l *Linked) Prepend(n *NoPrintln()

// 	second := l.Head

// 	l.Head = n
// 	l.Head.Next = second
// 	l.Lenght++
// }

// //Print the linkedList  data

// func (l Linked) PrinDataOnNode() {

// 	toPrint := l.Head
// 	for l.Lenght != 0 {
// 		fmt.Printf("%d ", toPrint.Data)
// 		toPrint = toPrint.Next
// 		l.Lenght--

// 	}
// 	fmt.Printf("\n")
// }

// //deleting the value from the linked List

// func (l *Linked) DeleteLinkedListData(value int) {

// 	previousToDelete := l.Head

// 	if l.Head == nil {
// 		return

// 	}
// 	if l.Head.Data == value {
// 		l.Head = l.Head.Next
// 		l.Lenght--
// 		return
// 	}

// 	for previousToDelete.Next.Data != value {

// 		if previousToDelete.Next.Next == nil {
// 			fmt.Println("there is no data for delete")
// 			return
// 		}
// 		previousToDelete = previousToDelete.Next
// 	}

// 	previousToDelete.Next = previousToDelete.Next.Next
// 	l.Lenght--
// }
