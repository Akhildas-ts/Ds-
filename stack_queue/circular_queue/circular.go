package main

import "fmt"


// front for get the frist element 
// rear means the current element 
// capacity is the total size of the CQ
// size  means how many elements availble ,so that size it means

type MyCircularQueue struct {
    size     int
    capacity int
    front    int
    rear     int
    data     []int
}

// constructor for calling CQ 

func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        size:     0,
        capacity: k,
        front:    0,
        rear:     -1,
        data:     make([]int, k),
    }
}

// Adding value on the CQ 

func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    this.rear = (this.rear + 1) % this.capacity
    this.data[this.rear] = value
    this.size++
    return true
}

// Deleteing Value in the  CQ 

func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.front = (this.front + 1) % this.capacity
    this.size--
    return true
}

// Getting the Frist Value 

func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.front]
}

// Getting the current value

func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.rear]
}

// Cheching is that CQ is empty or not 

func (this *MyCircularQueue) IsEmpty() bool {
    return this.size == 0
}


// Cheching is the CQ is full or not 

func (this *MyCircularQueue) IsFull() bool {
    return this.size == this.capacity
}


func main() {
	circularQueue := Constructor(8)

	// Enqueue elements into the circular queue
	fmt.Println("Enqueue 1:", circularQueue.EnQueue(1)) // Output: true
	fmt.Println("Enqueue 2:", circularQueue.EnQueue(2)) // Output: true
	fmt.Println("Enqueue 3:", circularQueue.EnQueue(3)) // Output: true
	fmt.Println("Enqueue 4:", circularQueue.EnQueue(4)) // Output: false (Queue is full)

	// Display front and rear elements
	fmt.Println("Front:", circularQueue.Front()) // Output:
	fmt.Println("last ",circularQueue.Rear())
}
