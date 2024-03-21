package main

import "fmt"

type circularLinked struct {
	head *node
}

type node struct {
	data int
	next *node
}

func (cl *circularLinked) insert(x int) {

	addNode := &node{data: x}

	if cl.head == nil {

		addNode.next = addNode
		cl.head = addNode

	} else {

		addNode.next = cl.head.next
		cl.head.next = addNode

	}
}

func (cl circularLinked) display() {

	if cl.head == nil {
        fmt.Println("Circular Linked List is empty")
        return
    }

    current := cl.head
    for {
        fmt.Printf("%d -> ", current.data)
        current = current.next
        if current == cl.head {
            break
        }
    }
    fmt.Println()
}


func main() {

	c1 := &circularLinked{}

	c1.insert(10)
	c1.insert(20)
	c1.insert(30)
	c1.insert(40)
	c1.insert(50)

	c1.display()
}
