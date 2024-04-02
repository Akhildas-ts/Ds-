package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// check the value is lessthan left then go to the left side of tree
// same as if value is greeter than go to the right side of tree
// when the node is nil that was the postion for value insert ..

func Insert(root *TreeNode, value int) *TreeNode {

	//here we got the insert position for the value
	if root == nil {
		return &TreeNode{val: value, left: nil, right: nil}
	}

	if value < root.val {
		//root will move the left side
		root.left = Insert(root.left, value)
	} else if value > root.val {
		// root will move the right side
		root.right = Insert(root.right, value)
	}

	return root

}

// we going to the smallest element at left
// then we start to print then move to right
func inorderTraversal(root *TreeNode) {

	if root != nil {
		inorderTraversal(root.left)
		fmt.Printf("%d ", root.val)
		inorderTraversal(root.right)
	}
}

func main() {
	var root *TreeNode
	value := []int{2, 5, 4, 5, 1, 3}

	for _, val := range value {
		
		root = Insert(root, val)
	}


	//    Delation(root,4)

	// fmt.Println(Search(root, 2))4
	

	fmt.Println("in order Traversal of the binary search  trer: ")
	inorderTraversal(root)
	// fmt.Println()

}

// search node into the tree

func Search(root *TreeNode, val int) bool {

	if root == nil {

		return false
	}

	if root.val == val {

		return true
	}

	if val < root.val {

		Search(root.left, val)
	} else if val > root.val {

		Search(root.right, val)
	}

	return false

}

//3.  delete Node Into a Tree


// we need to find the value  in these case, of senerio


// algoritham for deletion 	
// frist we need to reach the reach the deleation node 
// check if the left or right is nil (then it was easy to delete ) ,if its nil assign current node = left or right
// or if not nil , there  have 2 children nodes then find the right smallest element 
// assign the element to current root val  .
// then remove the last right value of delete root . 


func Delation(root  *TreeNode,value int )*TreeNode{

	if root == nil{
		return nil
	}


	if value < root.val{

		root.left = Delation(root.left,value)
	} else if value >root.val{
		root.right = Delation(root.right,value)
	}else{

       

		if root.left == nil{
			return root.right
		}else if root.right == nil{
			return root.left
		}


       // Right small index of value
	    Rsmall := FindMin(root.right)
		root.val = Rsmall.val 

		root.right = Delation(root.right,Rsmall.val)




	}

	return root



}

func FindMin(node *TreeNode)*TreeNode{


	for node.left != nil{

		node = node.left
	}
	return node
		
}