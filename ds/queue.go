package ds
import "fmt"

type Queue struct {
	list LinkedList
}

func (q *Queue) Push(value string) {
	q.list.Insert(value)
}

func (q *Queue) Pop() (string, error) {
	if q.list.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}
	msg := q.list.Head.data
	q.list.RemoveAt(0)
	return msg, nil
}