package main

import "fmt"

type MyCircularDeque struct {
	size     int
	capacity int
	front    int
	rare     int
	queue1   []int
}

func Constructor(k int) MyCircularDeque {

	return MyCircularDeque{
		size:     0,
		capacity: k,
		front:    0,
		rare:     -1,
		queue1:   make([]int, k),
	}

}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	// Shift elements to the right to make space for the new value

	this.front = (this.front - 1 + this.capacity) % this.capacity
	this.queue1[this.front] = value
	this.size++
	this.size++
	return true

}

func (this *MyCircularDeque) InsertLast(value int) bool {

	if this.IsFull() {
		return false
	}

	this.rare = (this.rare + 1) % this.capacity

	this.queue1[this.rare] = value
	this.size++
	return true

}

func (this *MyCircularDeque) DeleteFront() bool {

	if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % this.capacity
	this.size--
	return false

}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.rare = (this.rare + 1) % this.capacity
	this.size--
	return false

}

func (this *MyCircularDeque) GetFront() int {

	if this.IsEmpty() {
		return -1
	}

	return this.queue1[this.front]

}

func (this *MyCircularDeque) GetRear() int {

	if this.IsEmpty() {
		return -1
	}
	return this.queue1[this.rare]

}

func (this *MyCircularDeque) IsEmpty() bool {

	return this.size == 0

}

func (this *MyCircularDeque) IsFull() bool {

	return this.size == this.capacity

}

func main() {

	ob := Constructor(7)
	fmt.Println(ob.InsertFront(10))
	fmt.Println(ob.InsertLast(20))
	fmt.Println(ob.InsertLast(30))
	fmt.Println(ob.InsertLast(40))
	// fmt.Println(ob.InsertLast(2))
	// fmt.Println(ob.InsertLast(3))
	fmt.Println(ob.GetFront())
	fmt.Println(ob.GetRear())
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
