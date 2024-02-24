package Queue
// Queue that represent a queue that hold the a slice

type Queue struct {
	items []int
}

//Queue

func (q *Queue) Queue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	removedElement := q.items[0]
	q.items = q.items[1:]
	return removedElement
}
