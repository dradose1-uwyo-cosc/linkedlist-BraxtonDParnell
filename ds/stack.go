package ds

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(value string) { //Stack is first in last out, so push will just insert at the end of the list.
	s.list.Insert(value)
}

func (s *Stack) Pop() (string, bool) { //Remove the last element in the stack so use LinkedList's RemoveAt(x) function where x is the end of the list.
	if s.list.IsEmpty() {
		return "Stack is already empty", false
	}
	msg := s.list.Tail.data
	s.list.RemoveAt(s.list.GetSize() - 1)
	return msg, true
}
