package main

import (
	"errors"
	"fmt"
)

type QueueWithStacks[T any] struct {
	stack1 []T
	stack2 []T
}

// NewQueue creates a new queue using two stacks
func NewQueue[T any]() *QueueWithStacks[T] {
	return &QueueWithStacks[T]{}
}

// Enqueue adds an element to the queue
func (q *QueueWithStacks[T]) Enqueue(value T) {
	q.stack1 = append(q.stack1, value)
}

// Dequeue removes and returns the front element of the queue
func (q *QueueWithStacks[T]) Dequeue() (T, error) {
	// If both stacks are empty, return an error
	if len(q.stack1) == 0 && len(q.stack2) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}

	// If stack2 is empty, move elements from stack1 to stack2
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			// Pop from stack1 and push to stack2
			q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
			q.stack1 = q.stack1[:len(q.stack1)-1]
		}
	}

	// Pop from stack2 (this is the front of the queue)
	value := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return value, nil
}

// Front returns the front element of the queue without removing it
func (q *QueueWithStacks[T]) Front() (T, error) {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
			q.stack1 = q.stack1[:len(q.stack1)-1]
		}
	}

	if len(q.stack2) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}

	return q.stack2[len(q.stack2)-1], nil
}

// Size returns the number of elements in the queue
func (q *QueueWithStacks[T]) Size() int {
	return len(q.stack1) + len(q.stack2)
}

func main() {
	// Create a new queue of integers
	queue := NewQueue[int]()

	// Enqueue elements
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
	}

	// Print the front element
	if front, err := queue.Front(); err == nil {
		fmt.Println("Front element:", front)
	}

	// Dequeue elements
	for i := 1; i <= 3; i++ {
		if value, err := queue.Dequeue(); err == nil {
			fmt.Println("Dequeued:", value)
		}
	}

	// Enqueue more elements
	queue.Enqueue(6)
	queue.Enqueue(7)

	// Print the front element after enqueueing more elements
	if front, err := queue.Front(); err == nil {
		fmt.Println("Front element after enqueueing more:", front)
	}

	// Check size of the queue
	fmt.Println("Queue size:", queue.Size())
}
