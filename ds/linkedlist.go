package ds

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Insert(value string) { //Inserts at the end of the list.
	newNode := &Node{data: value}

	if l.Head == nil { //This is for if the Linked List was empty.
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}
	//Else just insert normally.
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Size++
}

func (l *LinkedList) InsertAt(position int, value string) error {
	newNode := &Node{data: value}

	if position == 0 { //Special case if the desired position is at 0
		newNode.Next = l.Head
		l.Head = newNode
		if l.Tail == nil { //If the ll was empty.
			l.Tail = newNode
		}
		l.Size++
		return nil
	}
	if position < 0 { //Checks for less than zero position input
		return fmt.Errorf("Position is below zero")
	}
	//Else, iterate through the ll to find the desired position
	current := l.Head
	for i := 0; i < position-1; i++ {
		if current == nil { //Error for if the position was out of bounds.
			return fmt.Errorf("Position is out of bounds.")
		}
		current = current.Next
	}
	//Getting here means we have found the position and we can insert the value
	newNode.Next = current.Next
	current.Next = newNode

	if newNode.Next == nil { //If the position is at the end of the ll.
		l.Tail = newNode
	}
	l.Size++
	return nil

}

func (l *LinkedList) Remove(value string) error {
	if l.Head == nil { //Error for empty list
		return fmt.Errorf("Can't remove: Empty list")
	}

	if l.Head.data == value { //Base case for if the first encountered node is the desired value to remove
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return nil
	}
	//Iterating through the LL to find the value and remove it.
	prev := l.Head
	current := l.Head.Next
	for current != nil {
		if current.data == value {
			prev.Next = current.Next
			if current.Next == nil { //If the removed value was the tail, make sure tail is updated.
				l.Tail = prev
			}
			return nil
		}
		prev = current
		current = current.Next
	}
	return fmt.Errorf("Can't remove: Value not found") //If the value was not found.
}

func (l *LinkedList) RemoveAll(value string) error {
	for l.Head != nil && l.Head.data == value { //Special case if there are one or multiple occurences of value at the front.
		l.Size--
		l.Head = l.Head.Next
	}
	if l.Head == nil { //If the list is empty.
		l.Tail = nil
		return nil
	}

	//Iterating through the list to find each occurence of value, removing it, and making sure no nodes are skipped.
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
	if l.IsEmpty() { //Empty list
		return fmt.Errorf("Can't remove: List is empty")
	}

	if pos == 0 { //Special case for if the position is the head.
		l.Head = l.Head.Next
		if l.Head == nil { //If the list is now empty.
			l.Tail = nil
		}
		l.Size--
		return nil
	}

	if pos < 0 { //If given position is less than 0.
		return fmt.Errorf("Position is less than 0.")
	}

	current := l.Head.Next
	prev := l.Head
	index := 1
	//Else, iterating through the list to get to the correct position.
	for current != nil && index <= pos {
		if index == pos {
			prev.Next = current.Next
			if current.Next == nil { //If the position was the tail, make previous the new tail.
				l.Tail = prev
			}
			l.Size--
			return nil
		}
		prev = current
		current = current.Next
		index++
	}
	return fmt.Errorf("Can't remove: position not in list") //Error if the position was out of bounds.
}

func (l *LinkedList) IsEmpty() bool { //Checks if list is empty.
	return l.Head == nil
}

func (l *LinkedList) GetSize() int { //Returns the size of the list.
	return l.Size
}

func (l *LinkedList) Reverse() {
	var prev *Node // A node with Data = nil and Next = nil
	current := l.Head
	l.Tail = l.Head // Since we are reversing, the current head will become the new tail.

	for current != nil {
		next := current.Next // Saving the next node for the current node
		current.Next = prev  // Reversing
		prev = current       // Moving forward by one node.
		current = next
	}

	l.Head = prev // The new head should now be the old tail.
}

func (l *LinkedList) PrintList() { //Prints out the nodes of the linked list.
	current := l.Head
	fmt.Print("Linked List: ")
	for current != nil {
		fmt.Print(current.data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}
