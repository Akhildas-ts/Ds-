package problem

import "fmt"

// func ContainsDuplicate(nums []int) bool {

// 	store := make(map[int]bool)

// 	for _, val := range nums {

// 		if store[val] {
// 			fmt.Println("true")
// 			return true
// 		} else {

// 			store[val] = true
// 		}

// 	}

// 	fmt.Println("false ")
// 	return false

// }

// * so these was the function , we need to write without the map, its pure hash table, //

type HashTable struct {
	data map[int]bool
}

func NewHashTable() *HashTable {
	return &HashTable{
		data: make(map[int]bool),
	}
}

func (ht *HashTable) Insert(val int) {
	ht.data[val] = true
}

func (ht *HashTable) Contains(val int) bool {
	return ht.data[val]
}

func ContainsDuplicate(nums []int) bool {
	ht := NewHashTable()

	for _, val := range nums {

		if ht.Contains(val) {
			fmt.Println("true ")
			return true
		} else {
			ht.Insert(val)
		}
	}

	fmt.Println("false ")

	return false
}
