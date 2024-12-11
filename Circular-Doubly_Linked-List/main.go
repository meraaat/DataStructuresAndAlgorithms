package main

import "fmt"

// Node represents a node in the circular doubly linked list
type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

// CircularDoublyLinkedList represents the circular doubly linked list
type CircularDoublyLinkedList struct {
	Tail *Node // Tail points to the last node in the list
}

// AddToEmpty adds a node to an empty list
func (cdll *CircularDoublyLinkedList) AddToEmpty(value int) {
	if cdll.Tail != nil {
		fmt.Println("List is not empty!")
		return
	}

	newNode := &Node{Value: value}
	newNode.Next = newNode // Points to itself
	newNode.Prev = newNode // Points to itself
	cdll.Tail = newNode
}

// AddToEnd adds a node at the end of the circular doubly linked list
func (cdll *CircularDoublyLinkedList) AddToEnd(value int) {
	if cdll.Tail == nil {
		cdll.AddToEmpty(value)
		return
	}

	newNode := &Node{Value: value}
	head := cdll.Tail.Next // Head node of the list

	newNode.Next = head   // New node points to the head
	newNode.Prev = cdll.Tail // New node points to the old tail
	cdll.Tail.Next = newNode // Old tail points to the new node
	head.Prev = newNode  // Head node points back to the new node
	cdll.Tail = newNode  // Update tail to the new node
}

// AddToBeginning adds a node at the beginning of the circular doubly linked list
func (cdll *CircularDoublyLinkedList) AddToBeginning(value int) {
	if cdll.Tail == nil {
		cdll.AddToEmpty(value)
		return
	}

	newNode := &Node{Value: value}
	head := cdll.Tail.Next // Current head node

	newNode.Next = head   // New node points to the old head
	newNode.Prev = cdll.Tail // New node points to the tail
	head.Prev = newNode   // Old head points back to the new node
	cdll.Tail.Next = newNode // Tail points to the new node
}

// Delete deletes the first occurrence of a node with the given value
func (cdll *CircularDoublyLinkedList) Delete(value int) {
	if cdll.Tail == nil {
		fmt.Println("List is empty!")
		return
	}

	// Traverse the list to find the node
	current := cdll.Tail.Next // Start at the head node
	for current.Value != value {
		if current == cdll.Tail { // Completed a full circle
			fmt.Println("Value not found in the list!")
			return
		}
		current = current.Next
	}

	// Single-node list case
	if current == cdll.Tail && current.Next == cdll.Tail {
		cdll.Tail = nil
		return
	}

	// Node to delete is the tail
	if current == cdll.Tail {
		cdll.Tail = cdll.Tail.Prev
	}

	// Adjust the pointers to remove the node
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
}

// Display prints all the elements in the circular doubly linked list
func (cdll *CircularDoublyLinkedList) Display() {
	if cdll.Tail == nil {
		fmt.Println("List is empty!")
		return
	}

	current := cdll.Tail.Next // Start at the head
	for {
		fmt.Print(current.Value, " ")
		current = current.Next
		if current == cdll.Tail.Next { // Completed a full cycle
			break
		}
	}
	fmt.Println()
}

// DisplayReverse prints all the elements in reverse order
func (cdll *CircularDoublyLinkedList) DisplayReverse() {
	if cdll.Tail == nil {
		fmt.Println("List is empty!")
		return
	}

	current := cdll.Tail // Start at the tail
	for {
		fmt.Print(current.Value, " ")
		current = current.Prev
		if current == cdll.Tail { // Completed a full cycle
			break
		}
	}
	fmt.Println()
}

// Main function to demonstrate the functionalities
func main() {
	cdll := &CircularDoublyLinkedList{}

	// Add nodes
	cdll.AddToEmpty(10)
	cdll.AddToEnd(20)
	cdll.AddToEnd(30)
	cdll.AddToBeginning(5)

	// Display list
	fmt.Println("Circular Doubly Linked List (Forward):")
	cdll.Display()

	// Display list in reverse
	fmt.Println("Circular Doubly Linked List (Reverse):")
	cdll.DisplayReverse()

	// Delete nodes
	cdll.Delete(10)
	fmt.Println("After deleting 10 (Forward):")
	cdll.Display()

	cdll.Delete(30)
	fmt.Println("After deleting 30 (Forward):")
	cdll.Display()

	cdll.Delete(5)
	fmt.Println("After deleting 5 (Forward):")
	cdll.Display()

	cdll.Delete(20)
	fmt.Println("After deleting 20 (Forward):")
	cdll.Display()
}
