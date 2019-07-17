package queue

// An FIFO queue.
type Queue []int

// Pushes the element into the queue.
func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

//Pops element from head
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Returns if the queue is emptr or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
