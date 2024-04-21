package LinkList

import (
	"fmt"
)

// SORT THE LINKED LIST

func (l *Linked) SortList(head *Node) *Node {
	if head == nil {
		// Nothing to sort if the list is empty or has only one node
		return nil
	}

	if head != nil && head == nil {
		return head
	}

	// Iterate through the list using bubble sort algorithm
	for current := head; current != nil; current = current.Next {
		fmt.Println("printing the currentnode", current.Data)
		for nextNode := current.Next; nextNode != nil; nextNode = nextNode.Next {
			fmt.Println("prini the next node", nextNode.Data)
			if current.Data > nextNode.Data {
				// Swap the data of the current node and the next node
				current.Data, nextNode.Data = nextNode.Data, current.Data
			}
		}
	}

	// Print the sorted list
	l.NodePrint(head)
	println()
	return head
}
