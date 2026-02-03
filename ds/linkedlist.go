package ds
import "fmt"
type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Insert(value string){
	newNode := &Node{data: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Size++
}

func (l *LinkedList) InsertAt(position int, value string) error {
	newNode := &Node{data: value}

	if position == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		if l.Tail == nil {
			l.Tail = newNode
		}
		l.Size++ 
		return nil
	}

	current := l.Head
	index := 0
	for current != nil && index < position-1 {
		current = current.Next
		index++
	}

	if current == nil {
		return fmt.Errorf("position out of bounds")
	}

	newNode.Next = current.Next
	current.Next = newNode

	if (newNode.Next == nil) {
		l.Tail = newNode
	}
	l.Size++
	return nil

}

func (l *LinkedList) Remove(value string) error {
	if l.Head == nil {
		return fmt.Errorf("Can't remove: Empty list")
	}

	if l.Head.data == value {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return nil
	}
	prev := l.Head
	current := l.Head.Next
	for current != nil {
		if current.data == value {
			prev.Next = current.Next
			if current.Next == nil {
				l.Tail = prev
			}
			return nil
		}
		prev = current
		current = current.Next
	}
	return fmt.Errorf("Can't remove: Value not found")
}

func (l *LinkedList) RemoveAll(value string) error {
	for l.Head != nil && l.Head.data == value {
		l.Size--
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return nil
	}
	prev := l.Head
	current := l.Head.Next
	for current != nil {
		if current.data == value {
			prev.Next = current.Next
			if current.Next == nil {
				l.Tail = prev
			}
			l.Size--
			current = prev.Next
		} else {
		prev = current
		current = current.Next
		}
	}
	return nil
}

func (l *LinkedList) RemoveAt(pos int) error {
	if l.IsEmpty(){
		return fmt.Errorf("Can't remove: List is empty")
	}

	if pos == 0 {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		l.Size--
		return nil
	}

	current := l.Head.Next
	prev := l.Head
	index := 1

	for current != nil && index <= pos {
		if index == pos {
			prev.Next = current.Next
			if current.Next == nil {
				l.Tail = prev
			}
			l.Size--
			return nil
		}
		prev = current
		current = current.Next
		index++
	}
	return fmt.Errorf("Can't remove: position not in list")
}

func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkedList) GetSize() int {
	return l.Size
}

func (l *LinkedList) Reverse() {
	var prev *Node
    current := l.Head
    l.Tail = l.Head // old head will become the new tail

    for current != nil {
        next := current.Next   // save next
        current.Next = prev    // reverse the link
        prev = current         // move prev forward
        current = next         // move current forward
    }

    l.Head = prev // new head is the old tail
}

func (l *LinkedList) PrintList() {
	current := l.Head
	fmt.Print("LinkedList: ")
	for current != nil {
		fmt.Print(current.data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}
