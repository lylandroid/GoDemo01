package main

import (
	".."
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop(), q.Pop(), q.IsEmpty(), q.Pop(), q.IsEmpty(), "---------------------")
}
