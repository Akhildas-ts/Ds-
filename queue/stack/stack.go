package main

import (
	"fmt"
)

type Mystack struct {
	queue1 []int
	queue2 []int
}

func Construct() Mystack {

	return Mystack{}
}

func (this *Mystack) Push(x int) {

	this.queue2 = append(this.queue2, x)

	for len(this.queue1) > 0 {

		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *Mystack) Pop() int {

	if len(this.queue1) == 0 {
		return -1
	}
	top := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return top
}

func (this *Mystack) Top() int {
	if len(this.queue1) == 0 {
		return -1
	}

	top := this.queue1[0]
	return top

}

func (this *Mystack) Empty() bool {
	return len(this.queue1) == 0
}

func main() {

	ob := Construct()
	ob.Push(1)
	ob.Push(2)
	fmt.Println(ob.Top())
	fmt.Println(ob.Pop())
	fmt.Println(ob.Empty())

}
