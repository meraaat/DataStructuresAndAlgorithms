package main

import "fmt"

// Node represents a node in the circular singly linked list
type Node struct {
	Value int
	Next  *Node
}

// CircularSinglyLinkedList represents the circular singly linked list
type CircularSinglyLinkedList struct {
	Tail *Node // Tail points to the last node in the list
}

// AddToEmpty adds a node to an empty list
func (csll *CircularSinglyLinkedList) AddToEmpty(value int) {
	if csll.Tail != nil {
		fmt.Println("List is not empty!")
		return
	}

	newNode := &Node{Value: value}
	newNode.Next = newNode // Point to itself
	csll.Tail = newNode
}

// AddToEnd adds a node at the end of the circular singly linked list
func (csll *CircularSinglyLinkedList) AddToEnd(value int) {
	if csll.Tail == nil {
		csll.AddToEmpty(value)
		return
	}

	newNode := &Node{Value: value}
	newNode.Next = csll.Tail.Next // New node points to the head
	csll.Tail.Next = newNode      // Old tail points to the new node
	csll.Tail = newNode           // Update tail
}

// AddToBeginning adds a node at the beginning of the list
func (csll *CircularSinglyLinkedList) AddToBeginning(value int) {
	if csll.Tail == nil {
		csll.AddToEmpty(value)
		return
	}

	newNode := &Node{Value: value}
	newNode.Next = csll.Tail.Next // New node points to the head
	csll.Tail.Next = newNode      // Tail points to the new node
}

// Delete deletes the first occurrence of a node with the given value
func (csll *CircularSinglyLinkedList) Delete(value int) {
	if csll.Tail == nil {
		fmt.Println("List is empty!")
		return
	}

	// Check if the list has only one node
	if csll.Tail.Next == csll.Tail && csll.Tail.Value == value {
		csll.Tail = nil
		return
	}

	// Traverse the list to find the node
	current := csll.Tail.Next
	prev := csll.Tail

	for current.Value != value {
		if current == csll.Tail { // Value not found
			fmt.Println("Value not found in the list!")
			return
		}
		prev = current
		current = current.Next
	}

	// Node to delete is found
	if current == csll.Tail { // If the node is the tail
		csll.Tail = prev
	}
	prev.Next = current.Next // Remove the node from the list
}

// Display prints all the elements in the circular singly linked list
func (csll *CircularSinglyLinkedList) Display() {
	if csll.Tail == nil {
		fmt.Println("List is empty!")
		return
	}

	current := csll.Tail.Next // Start from the head
	for {
		fmt.Print(current.Value, " ")
		current = current.Next
		if current == csll.Tail.Next { // Completed a full cycle
			break
		}
	}
	fmt.Println()
}

// Main function to demonstrate the functionalities
func main() {
	csll := &CircularSinglyLinkedList{}

	// Add nodes
	csll.AddToEmpty(10)
	csll.AddToEnd(20)
	csll.AddToEnd(30)
	csll.AddToBeginning(5)

	// Display list
	fmt.Println("Circular Singly Linked List:")
	csll.Display()

	// Delete nodes
	csll.Delete(10)
	fmt.Println("After deleting 10:")
	csll.Display()

	csll.Delete(30)
	fmt.Println("After deleting 30:")
	csll.Display()

	csll.Delete(5)
	fmt.Println("After deleting 5:")
	csll.Display()

	csll.Delete(20)
	fmt.Println("After deleting 20:")
	csll.Display()
}
