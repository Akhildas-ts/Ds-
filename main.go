package main

import (
	Linklist "stack/linkedlist"
)

// Stack
type Stack struct {
	items []int
}

// push element into the Stack
func (s *Stack) Push(i int) {

	s.items = append(s.items, i)

}

//Pop will remove the value at the end and
// return the removed value

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	removedElement := s.items[l]
	s.items = s.items[:l]
	return removedElement
}
func main() {

	//stack

	// mystack := Stack{}
	// fmt.Println(mystack)
	// mystack.Push(100)
	// mystack.Push(200)
	// mystack.Push(300)
	// fmt.Println(mystack)
	// mystack.Pop()
	// fmt.Println(mystack)

	//queue and deque
	// q := Queue.Queue{}

	// q.Queue(2)
	// q.Queue(20)
	// q.Queue(40)
	// fmt.Println(q)

	// Linked list

	node1 := &Linklist.Node{Data: 10}
	node2 := &Linklist.Node{Data: 20}
	node3 := &Linklist.Node{Data: 30}
	node4 := &Linklist.Node{Data: 40}
	node5 := &Linklist.Node{Data: 50}
	node6 := &Linklist.Node{Data: 60}
	node7 := &Linklist.Node{Data: 70}
	linkedList := Linklist.Linked{}

	linkedList.Prepend(node1)
	linkedList.Prepend(node2)
	linkedList.Prepend(node3)
	linkedList.Prepend(node4)
	linkedList.Prepend(node5)
	linkedList.Prepend(node6)
	linkedList.Prepend(node7)

	linkedList.PrinDataOnNode()

	// fmt.Println(linkedList)

	linkedList.DeleteLinkedListData(70)
	linkedList.PrinDataOnNode()
	// fmt.Println("LinkedList after delete",linkedList)

	// fmt.Println("Length of linked list:", linkedList.Lenght)
	// fmt.Println("Data of the first node:", linkedList.Head.Data)

}
