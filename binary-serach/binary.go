package binaryserach

//Node represent the componets of a binary serach tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert  will Add a node to the tree
//The key to add should not be already in the tree

func (n *Node) Insert(k int) {

	if n.Key < k {
		//move right

		if n.Right == nil {

			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}

	} else if n.Key > k {
		//move left

		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}

	}

}

//Search will take in a key value
// return the value with node

func (n *Node) Serach(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Serach(k)

		
	} else if n.Key > k {
		
	return n.Left.Serach(k)

	}

	return true

}
