//Braxton Parnell
//COSC 3750
//02/02/2026
//
/*
	Don't forget to run your go mod init command in your terminal
	Review the assignment instructions for running your code
	All the code you need to write should be put in the /ds/ package files
	Uncomment the import statement for your module name
	you can uncomment the tests in main as you go to test
	The code in main is not an extensive test, you should add more and test your code as needed
*/
package main

import (
	"fmt"
	"linkedlist-BraxtonDParnell/ds"  //this needs to be changes to your module name and uncommented
)

func main() {

	linkedlist := &ds.LinkedList{}
	linkedlist.InsertAt(0, "first")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")
	linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	stack := &ds.Stack{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	data, _ := stack.Pop()
	println("Popped from stack:", data)

	queue := &ds.Queue{}
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")
	data, _ = queue.Pop()
	println("Popped from queue:", data)
}
