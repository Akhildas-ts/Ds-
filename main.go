package main

import (
	"fmt"
	binaryserach "stack/binary-serach"
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

	//STACK

	// mystack := Stack{}
	// fmt.Println(mystack)
	// mystack.Push(100)
	// mystack.Push(200)
	// mystack.Push(300)
	// fmt.Println(mystack)
	// mystack.Pop()
	// fmt.Println(mystack)

	//QUEUE AND DEQUEUE

	// q := Queue.Queue{}

	// q.Queue(2)
	// q.Queue(20)
	// q.Queue(40)
	// fmt.Println(q)

	// LINKED LIST

	// node1 := &Linklist.Node{Data: 10}
	// node2 := &Linklist.Node{Data: 20}
	// node3 := &Linklist.Node{Data: 30}
	// node4 := &Linklist.Node{Data: 40}
	// node5 := &Linklist.Node{Data: 50}
	// node6 := &Linklist.Node{Data: 60}
	// node7 := &Linklist.Node{Data: 70}
	// linkedList := Linklist.Linked{}
	// linkedList.Prepend(node1)
	// linkedList.Prepend(node2)
	// linkedList.Prepend(node3)
	// linkedList.Prepend(node4)
	// linkedList.Prepend(node5)
	// linkedList.Prepend(node6)
	// linkedList.Prepend(node7)

	// linkedList.PrinDataOnNode()

	// linkedList.DeleteLinkedListData(70)
	// linkedList.PrinDataOnNode()

	//BINARY SERACH

	tree := &binaryserach.Node{Key: 39}

	tree.Insert(400)
	tree.Insert(40)
	tree.Insert(80)
	tree.Insert(4)
	tree.Insert(30)
	tree.Insert(90)

	fmt.Println(tree.Serach(39))

}
