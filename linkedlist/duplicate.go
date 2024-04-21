package Linklist

import "fmt"

//copy the node
// need a map to store
//if there is any duplicate value we found then change the direction

func (l *Linked) RemoveDuplicate() {

	c := l.Head

	store := make(map[int]bool)

	store[c.Data] = true

	for c.Next != nil {
		if store[c.Next.Data] {
			c.Next = c.Next.Next
		} else {
			store[c.Next.Data] = true
			c = c.Next
		}
	}
}
func DisplayList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
