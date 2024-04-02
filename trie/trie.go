package main

import "fmt"

const AlphabetSize = 26

// Node Represent each node in tree
type Node struct {
	children [AlphabetSize]*Node
}

//trie represenet a trie and has a pointer to the root node

type trie struct {
	root *Node
}

func Init() *trie {

	result := &trie{root: &Node{}}
	return result
}

func (t *trie) insertIntoTrie(s string) {

	currentIdx := t.root

	for _, value := range s {

		indx :=  value-'a' 


		if currentIdx.children[indx]  == nil{
			currentIdx.children[indx] = &Node{}
		}

		currentIdx = currentIdx.children[indx]

		

	}

}

func main() {

	tri := Init()
	tri.insertIntoTrie("akhil")
	fmt.Println(tri)

}
