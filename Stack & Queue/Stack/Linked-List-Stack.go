package main

import (
	"errors"
)

// Node represents a single element in the linked list
type Node[T any] struct {
	value T        // Generic value
	next  *Node[T] // Pointer to the next node
}

// StackG represents a generic Stack implemented using a linked list
type StackG[T any] struct {
	top  *Node[T] // Pointer to the top of the Stack
	size int      // Number of elements in the Stack
}

// Push adds an element to the top of the Stack
func (s *StackG[T]) Push(value T) {
	newNode := &Node[T]{value: value, next: s.top} // Create a new node pointing to the current top
	s.top = newNode                                // Update the top pointer
	s.size++                                       // Increment size
}

// Pop removes and returns the top element of the Stack
func (s *StackG[T]) Pop() (T, error) {
	var zeroValue T // Default zero value for type T
	if s.IsEmpty() {
		return zeroValue, errors.New("StackG is empty") // Handle underflow
	}
	value := s.top.value // Retrieve the value of the top node
	s.top = s.top.next   // Update the top pointer to the next node
	s.size--             // Decrement size
	return value, nil
}

// Peek returns the top element without removing it
func (s *StackG[T]) Peek() (T, error) {
	var zeroValue T // Default zero value for type T
	if s.IsEmpty() {
		return zeroValue, errors.New("StackG is empty") // Handle empty Stack
	}
	return s.top.value, nil
}

// IsEmpty checks if the StackG is empty
func (s *StackG[T]) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of elements in the StackG
func (s *StackG[T]) Size() int {
	return s.size
}
