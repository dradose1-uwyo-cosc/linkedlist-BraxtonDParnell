package ds
type Stack struct {
	list LinkedList
}

func (s *Stack) Push(value string) {
	s.list.Insert(value)
}

func (s *Stack) Pop() (string, bool) {
	if s.list.IsEmpty(){
		return "Stack is already empty", false
	}
	msg := s.list.Tail.data
	s.list.RemoveAt(s.list.GetSize()-1)
	s.list.PrintList()
	return msg, true
}