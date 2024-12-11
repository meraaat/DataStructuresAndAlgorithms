package main

import "fmt"

// Stack with generics
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, error) {

	if len(s.items) == 0 {
		var zero T
		return zero, fmt.Errorf("Stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, nil
}

// Peek returns the last item without removing it
func (s *Stack[T]) Peek() (T, error) {

	if len(s.items) == 0 {
		var zero T
		return zero, fmt.Errorf("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
