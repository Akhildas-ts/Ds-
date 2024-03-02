package main

import (
	"fmt"
	"strings"
)

type DNode struct {
	data int
	next *DNode
	prec *DNode
}

type DlinkedList struct {
	head *DNode
	tail *DNode
}

type Node struct {
	data int
	next *Node
}

type LinkedLiset struct {
	head   *Node
	length int
}

func (l *LinkedLiset) RemoveDuplicate() {

	store := make(map[int]bool)

	c := l.head
	store[c.data] = true

	for c.next != nil {

		if store[c.next.data] == true {
			c.next = c.next.next
		} else {
			store[c.next.data] = true
			c = c.next
		}

	}
}

func (l LinkedLiset) PrintNode() {
	forPrint := l.head

	for forPrint != nil {
		fmt.Printf(":%d ", forPrint.data)
		forPrint = forPrint.next
	}
}

func (l *LinkedLiset) InserValue(val int) {

	node1 := &Node{data: val}

	c := l.head
	if l.head == nil {
		l.head = node1
		return
	}

	for c.next != nil {
		c = c.next
	}
	c.next = node1
	l.length++

}
func (l *LinkedLiset) MiddleValue(val int) {

	addNOde := &Node{data: val}

	midpos := l.length / 2

	if l.head == nil {
		l.head = addNOde
		l.length++
		return
	}

	current := l.head
	for i := 1; i < midpos; i++ {
		current = current.next
	}

	addNOde.next = current.next
	current.next = addNOde
}

func (l *LinkedLiset) ArrayConverToLinked(arr []int) {

	for _, val := range arr {
		l.InserValue(val)

	}
}

func (l *LinkedLiset) DuplicateValue() {

}

func (dl *DlinkedList) AddElementDL(val int) {

	addNode := &DNode{data: val}

	if dl.head == nil {
		dl.head = addNode
		dl.tail = addNode

		return
	}

	addNode.prec = dl.tail
	dl.tail.next = addNode
	dl.tail = addNode
}

func (dl *DlinkedList) PrintDL() {

	forPrint := dl.head

	for forPrint != nil {
		fmt.Printf("%d ", forPrint.data)
		forPrint = forPrint.next
	}
}

func replaceAlphabets(input string, n int) string {
	// Define a constant representing the English alphabet
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	// Convert the input string to lowercase
	input = strings.ToLower(input)

	// Initialize an empty string to store the result
	result := ""

	// Iterate over each character in the input string
	for _, char := range input {
		// Check if the character is an alphabet
		if 'a' <= char && char <= 'z' {
			// Calculate the new index of the alphabet character
			newIndex := strings.IndexByte(alphabet, byte(char)) + n

			// Ensure the index stays within the bounds of the alphabet
			newIndex = (newIndex + 26) % 26

			// Append the new alphabet character to the result string
			result += string(alphabet[newIndex])
		} else {
			// If the character is not an alphabet, append it unchanged
			result += string(char)
		}
	}

	return result
}

func main() {

	l1 := &LinkedLiset{}
	l1.InserValue(10)
	l1.InserValue(20)
	l1.InserValue(10)
	// l1.InserValue(30)
	// l1.InserValue(20)
	// l1.InserValue(10)
	// l1.InserValue(20)

	// l1.PrintNode()
	l1.MiddleValue(1)
	println()
	// l1.PrintNode()
	// l1.RemoveDuplicate()
	// println()
	// l1.PrintNode()

	// dl := &DlinkedList{}
	// dl.AddElementDL(10)
	// dl.AddElementDL(20)
	// dl.AddElementDL(20)
	// dl.PrintDL()

	// arr := []int{1,3,5,1,3,5,5}
	// l1.ArrayConverToLinked(arr)
	// l1.PrintNode()

	st :=replaceAlphabets("afsgs",4)
	fmt.Println("st",st)

}l
















package main

import (
	"fmt"
)

type Node struct {
	data int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

type DNode struct {
	data int
	prev *DNode
	next *DNode
}

type DLinkedList struct {
	head   *DNode
	tail   *DNode
	length int
}

//

func (dl *DLinkedList) DLaddElement(val int) {

	addNode := &DNode{data: val}

	if dl.head == nil {
		dl.head = addNode
		dl.tail = addNode
		return
	}

	addNode.prev = dl.tail
	dl.tail.next = addNode
	dl.tail = addNode
	dl.length++
}

func (dl *DLinkedList)RemoveNTHelement(val int ){


	lastNode:= dl.tail

	for lastNode != nil{
		fmt.Println("current data ",lastNode.data)
		lastNode = lastNode.prec
	}

}

func (dl DLinkedList) DprintNode() {

	forPrint := dl.head

	for forPrint != nil {
		fmt.Printf("%d ", forPrint.data)
		forPrint = forPrint.next
	}
}

func (l *LinkedList) Duplicate() {

	store := make(map[int]bool)
	c := l.head

	if l.head == nil {
		fmt.Println("there is no value ")
		return
	}
	store[c.data] = true

	for c.Next != nil {
		if store[c.Next.data] {
			c.Next = c.Next.Next
		} else {
			store[c.Next.data] = true
			c = c.Next
		}

	}

}

func (l *LinkedList) Append(val int) {

	addNode := &Node{data: val}
	if l.head == nil {
		l.head = addNode
		return
	}

	c := l.head

	for c.Next != nil {
		c = c.Next
	}

	c.Next = addNode
}

func (l LinkedList) PrintNode() {
	forpRint := l.head

	for forpRint != nil {

		fmt.Printf("%d ", forpRint.data)

		forpRint = forpRint.Next

	}
}

func (l LinkedList) PrintReverse() {
	forHead := l.head

	l.PrintRecurtion(forHead)

}

func (l LinkedList) PrintRecurtion(node *Node) {

	if node == nil {
		return
	}
	l.PrintRecurtion(node.Next)
	fmt.Printf("%d ", node.data)

}

func (l *LinkedList) Delete(val int) {

	c := l.head

	if l.head.data == val && l.head.Next == nil {
		l.head = nil
		return

	}

	if l.head != nil && l.head.data == val {
		l.head = l.head.Next

	}

	for c.Next.data != val {
		c = c.Next

	}

	c.Next = c.Next.Next

}

func (l *LinkedList) InsertAfterX(x, val int) {

	addNode := &Node{data: val}

	c := l.head

	// if l.head.data == x{
	// 	addNode.Next = l.head.Next
	// 	l.head = addNode
	// 	return
	// }
	if l.head == nil {
		return
	}

	for c.data != x {

		c = c.Next

	}

	addNode.Next = c.Next
	c.Next = addNode
}

func (l *LinkedList) InsetBeforX(x, val int) {
	addNode := &Node{data: val}

	if l.head.data == val {
		addNode.Next = l.head
		return
	}
	c := l.head

	for c.Next.data != x {
		c = c.Next

	}
	addNode.Next = c.Next
	c.Next = addNode
}

func (dl *DLinkedList) DlmiddleElement(val int) {

	addNode := &DNode{data: val}

	midPos := dl.length / 2

	if dl.head == nil {
		dl.head = addNode
		dl.tail = addNode
		return
	}

	c := dl.head

	for i := 1; i < midPos; i++ {

		c = c.next
	}
	addNode.next = c.next

	if c.next != nil {
		c.next.prev = addNode
	}
	addNode.prev = c
	c.next = addNode
	dl.length++

}

func (l *LinkedList)ConvertarrayIntoLinkedList(arr []int){

	
	for _,val := range arr{
		l.AddElementToArray(val)
	
	}

	
}

func (l *LinkedList) AddElementToArray(val int){

	node1 := &Node{data: val}

	c:= l.head
	if l.head == nil{
		l.head = node1
		return
	}

	for c.Next != nil{
		c = c.Next
	}

	c.Next = node1

}



func main() {

		// l1 := &LinkedList{}

		// arr := []int{2,35,2,6,11}
		// l1.ConvertarrayIntoLinkedList(arr)
		// l1.PrintNode()

	// 	l1.Append(10)

	// 	l1.Append(20)
	// 	l1.Append(30)

	//   l1.InsetBeforX(20,39)
	// 	l1.PrintNode()
	// 	println()
	// 	l1.PrintNode()
	// l1.Append(10)
	// l1.Append(20)
	// l1.PrintNode()
	// println()
	// l1.PrintReverse()

	// DOUBULE LIST

	// dl := &DLinkedList{}

	// dl.DLaddElement(10)
	// dl.DLaddElement(20)
	// dl.DLaddElement(30)
	// dl.DLaddElement(40)
	// dl.DLaddElement(40)
	// dl.DprintNode()
	// dl.DlmiddleElement(29)
	// println()
	// dl.DprintNode()

}

