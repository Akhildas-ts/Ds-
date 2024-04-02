// package main

// import "golang.org/x/text/unicode/rangetable"

// const Arraysize = 10

// type HashTable struct {
// 	Array [Arraysize]*bucket
// }

// type bucket struct {
// 	head *bucketNode
// }

// type bucketNode struct {
// 	val  int
// 	next *bucketNode
// }

// // here we are inserting the node to the  single of array,

// func inti() *HashTable {

// 	result := &HashTable{}

// 	for i := range result.Array {

// 		result.Array[i] = &bucket{}
// 	}

// 	return result

// }

// // when we add a new element to the hash table , frist of all we need to hash the value of the element, then we go the postion ,
// // where we got a postion, where we need to insert

// func (b *bucket) Insert(k int) bool {

// }

// func hash(k []int) int {

// 	sum := 0

// 	for _, num := range k{
// 		sum += int(num)
// 	}

// 	}

