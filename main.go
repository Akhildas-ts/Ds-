package main

import "fmt"

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
	 l := len(s.items)-1
	removedElement := s.items[l]
	s.items = s.items[:l]
	return removedElement
}
func main() {

	mystack := Stack{}
	fmt.Println(mystack)
	mystack.Push(100)
	mystack.Push(200)
	mystack.Push(300)
	fmt.Println(mystack)
	mystack.Pop()
	fmt.Println(mystack)
}
