package double

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}  


type D_LinkedList struct {
	Head *Node
	Tail *Node
}


// Append a value into the Double linked list 

func (l *D_LinkedList) D_Append(val int) {

	node1 := &Node{Data: val}

	if l.Head == nil {
		l.Head = node1
		l.Tail = node1
		return
	}

	cuurentNode := l.Head

	for cuurentNode.Next != nil {

		cuurentNode = cuurentNode.Next
	}
	node1.Prev = l.Tail
	l.Tail.Next = node1
	l.Tail = node1

}

func (l D_LinkedList) PrintDouble() {

	forPrint := l.Head

	for forPrint != nil {

		fmt.Printf("%d ", forPrint.Data)
		forPrint = forPrint.Next

	}
}
