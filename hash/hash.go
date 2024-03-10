package hashs

import "fmt"

//array of size
const Arraysize = 7

// HashTable will Hold an array
type HashTable struct {
	Array [Arraysize]*Bucket
}

//bucket is linked list so we need to describe like that that

type Bucket struct {
	Head *bucketNode
}

//bucket node structure
type bucketNode struct {
	Key  string
	Next *bucketNode

	
}

//Init will create the bucket in the each slot of the array  of the hash table

func Init() *HashTable {

	result := &HashTable{}

	for i := range result.Array {

		result.Array[i] = &Bucket{}
	}
	return result

}

// hash

func Hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}
	return sum % Arraysize
}

//insert will take in a key, create a node with the key and insert the node in the bucket..


func (b *Bucket) Insert(k string) {

	newNode := &bucketNode{Key: k}

	newNode.Next = b.Head
	b.Head = newNode
}

//Serach for the key in the hash table in these case ..

func (b *Bucket) SerachHash(k string) bool {

	current := b.Head

	if b.Head == nil {
		return false
	}

	for current != nil {

		if current.Key == k {
			return true
		}
		current = current.Next
	}

	return false
}

 

// DELETIN THE ELEMENTS FROM TO THE HASH TABLE.. 

func (b *Bucket) DeleteHash(k string) {

	if b.Head == nil {
		fmt.Println("there is no nodes")
		return
	}

	if b.Head.Key == k {

		fmt.Println("value at head deleted")

		b.Head = b.Head.Next
		return
	}

	current := b.Head

	
	for current.Next != nil{

		if current.Next.Key == k{
			fmt.Println("value delteed ")
			return
		}

		current = current.Next

	}


}
