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
