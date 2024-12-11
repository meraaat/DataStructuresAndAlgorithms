package main

import "fmt"

func main() {

	list := &LinkedList{}

	list.AddToEnd(5)
	list.AddToEnd(10)
	list.AddToEnd(15)
	list.AddToEnd(20)

	list.AddToFront(25)
	list.AddToFront(30)

	list.Delete(10)
	list.Delete(30)
	list.Delete(28)

	list.Display()

	list.AddAfter(15, 16)
	list.Display()
}

// Node represent a single node in the singly link list
type Node struct {
	Value int
	Next  *Node
}

// LinkList represent the entire link list
type LinkedList struct {
	Head *Node
}

// AddToEnd adds a new node with the given value to the end of the list
func (l *LinkedList) AddToEnd(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// AddToFront adds a new node with the given value to the front of the list
func (l *LinkedList) AddToFront(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

// AddAfter adds a new node with the given value after the first occurrence of node with value x
func (l *LinkedList) AddAfter(x, value int) {

	if l.Head == nil {
		fmt.Println("List is empty. Cannot add.")
		return
	}

	current := l.Head
	for current != nil && current.Value != x {
		current = current.Next
	}

	if current == nil {
		fmt.Printf("value %d not found in the List. Cannot add %d.\n", x, value)
		return
	}

	newNode := &Node{Value: value, Next: current.Next}
	current.Next = newNode
	fmt.Printf("Added %d after %d.\n", value, x)

}

// Delete deletes the first node with the given value
func (l *LinkedList) Delete(value int) {

	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head

	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

// Display prints the values of the linked list
func (l *LinkedList) Display() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println(nil)
}
