package main

import "fmt"

type myque struct {
	queue1 []int
	queue2 []int
}

func Constructor() myque {
	return myque{}
}

func (this *myque) push(x int) {

	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	this.queue2 = append(this.queue2, x)

	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *myque) pop() int {
	top := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return top

}

func (this *myque) top() int {

	top := this.queue1[0]
	return top
}

func (this *myque)empty()bool{

	return len(this.queue1) == 0

}

func main() {

	ob := Constructor()

	ob.push(2)
	ob.push(3)
	ob.push(4)
	
	fmt.Println(ob.top())
	fmt.Println(ob.pop())
	fmt.Println(ob.empty())

}
