package Linklist

import "fmt"


// We add values on the node 


func (l *Linked) Append(value int) {
	
	node := &Node{Data: value, Next: nil}
	
	if l.Head == nil {
		l.Head = node
		return
	}
	
	lastNode := l.Head
	
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		
	}
	lastNode.Next = node
}


//array convert to linklist


func ConvertIntoArray(arr []int) *Linked {

	l1 := &Linked{}

	for _, value := range arr {
		l1.Append(value)
	}
	return l1

}
func (l *Linked) PrintListInArray() {
	currNode := l.Head
	for currNode != nil {
		fmt.Printf("%d -> ", currNode.Data)
		currNode = currNode.Next
	}
	fmt.Println("nil")
}
