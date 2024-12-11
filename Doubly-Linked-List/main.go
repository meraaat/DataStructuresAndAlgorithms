package main

import "fmt"

func main() {
	dll := &DoublyLinkedList{}

	dll.AddToEnd(10)
	dll.AddToEnd(20)
	dll.AddToFront(5)
	dll.AddToEnd(30)
	dll.DisplayForward()
	dll.DisplayBackward()

	dll.AddAfter(20, 25)
	dll.DisplayForward()

	dll.Delete(10)
	dll.DisplayForward()

	dll.Delete(5)
	dll.DisplayForward()

	dll.Delete(30)
	dll.DisplayForward()

}

// Node represents a node in the doubly linked list
type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// AddToEnd adds a new node with the given value at the end of the list
func (dll *DoublyLinkedList) AddToEnd(value int) {

	newNode := &Node{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode

}

// AddToFront adds a new node with the given value at the front of the list
func (dll *DoublyLinkedList) AddToFront(value int) {

	newNode := &Node{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	newNode.Next = dll.Head
	dll.Head.Prev = newNode
	dll.Head = newNode
}

// AddAfter inserts a new node with the given value after the first node with value x
func (dll *DoublyLinkedList) AddAfter(x, value int) {

	if dll.Head == nil {
		fmt.Println("List is empty. Cannot Add.")
		return
	}

	current := dll.Head
	for current != nil && current.Value != x {
		current = current.Next
	}

	if current == nil {
		fmt.Printf("value %d not found. Cannot add %d.\n", x, value)
		return
	}

	newNode := &Node{Value: value, Next: current.Next, Prev: current}

	if current.Next != nil {
		current.Next.Prev = newNode
	} else { // If current is the tail
		dll.Tail = newNode
	}

	current.Next = newNode
}

// Delete deletes the first occurrence of a node with the given value
func (dll *DoublyLinkedList) Delete(value int) {

	if dll.Head == nil {
		return
	}

	// Check if the head node matches the value
	if dll.Head.Value == value {
		dll.Head = dll.Head.Next
		if dll.Head != nil { // Update the Prev pointer for the new Head
			dll.Head.Prev = nil
		} else { // If the list is now empty, update the Tail
			dll.Tail = nil
		}
		return
	}

	// Check if the Tail node matches the value
	if dll.Tail.Value == value {
		dll.Tail = dll.Tail.Prev
		if dll.Tail != nil { // Update the Next pointer of the new Tail
			dll.Tail.Next = nil
		} else { // If the list is now empty, update the Head
			dll.Head = nil
		}
		return
	}

	// Traverse the list to find the node to delete
	current := dll.Head

	for current != nil && current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	// If the node is found, update pointers
	if current != nil && current.Next != nil {
		toDelete := current.Next
		current.Next = toDelete.Next
		if toDelete.Next != nil { // Update the Prev pointer of the next node
			toDelete.Next.Prev = current
		}
	}

}

// DisplayForward prints the values of the list from head to tail
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.Head

	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Next
	}
	fmt.Println(nil)
}

// DisplayBackward prints the values of the list from tail to head
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.Tail

	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Prev
	}
	fmt.Println(nil)
}
