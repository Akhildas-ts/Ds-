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


//insert before a element 
// if the head of node is that element we need to add new node after 
// other wise we need to find the x , so that time that time we are adding the before node into prev for add a new node 
//so when we get that x node, so that time before the x node are store in the prev, so we just add the node middle of (x and prev)

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

	fmt.Println("node  with data not found \n", x)

}


