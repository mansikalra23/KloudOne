package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println("Initial Queue : ", myQueue)
	myQueue.Enqueue(18)
	myQueue.Enqueue(78)
	myQueue.Enqueue(45)
	myQueue.Enqueue(34)
	myQueue.Enqueue(27)
	fmt.Println("After entering data : ", myQueue)
	myQueue.Dequeue()
	fmt.Println("After removing data : ", myQueue)

}
