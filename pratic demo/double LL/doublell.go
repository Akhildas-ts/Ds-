package double


type DNode struct{
	data int
	prev *DNode
	next *DNode
}

type DLinkedList struct{
	head *DNode
	tail *DNode
}


func(dl *DLinkedList)AddElement(val int ){

	addNode := &DNode{data: val}

	if dl.head == nil{

		dl.head = addNode
		dl.tail = addNode
		return
	}

	addNode.prev = dl.tail
	dl.tail.next = addNode
	dl.tail = addNode
}