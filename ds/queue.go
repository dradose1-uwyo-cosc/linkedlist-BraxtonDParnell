package ds

import "fmt"

type Queue struct {
	list LinkedList
}

func (q *Queue) Push(value string) { //Queues are first in first out, so when we push (enqueue) we just insert at the end of the LL like normal.
	q.list.Insert(value)
}

func (q *Queue) Pop() (string, error) { //When we pop (dequeue) we want to remove at the head instead of the tail compared to a stack. So use RemoveAt(0).
	if q.list.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}
	msg := q.list.Head.data
	q.list.RemoveAt(0)
	return msg, nil
}
