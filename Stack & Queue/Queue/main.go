package main

import (
	"errors"
	"fmt"
)

type CircularQueue[T any] struct {
	data                        []T
	front, rear, size, capacity int
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		data:     make([]T, capacity),
		front:    -1,
		rear:     -1,
		capacity: capacity,
	}
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *CircularQueue[T]) Enqueue(value T) error {
	if q.IsFull() {
		q.resize(q.capacity * 2) // Resize the queue to double the size when full
	}
	if q.front == -1 {
		q.front = 0
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = value
	q.size++
	return nil
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	if q.size == 0 {
		q.front = -1
		q.rear = -1
	}
	// Shrink the queue when it gets under 25% full
	if q.size <= q.capacity/4 && q.capacity > 1 {
		q.resize(q.capacity / 2)
	}
	return value, nil
}

func (q *CircularQueue[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	i := q.front
	for j := 0; j < q.size; j++ {
		newData[j] = q.data[i]
		i = (i + 1) % q.capacity
	}
	q.data = newData
	q.front = 0
	q.rear = q.size - 1
	q.capacity = newCapacity
}

func (q *CircularQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	return q.data[q.front], nil
}

func (q *CircularQueue[T]) Rear() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	return q.data[q.rear], nil
}

func (q *CircularQueue[T]) Size() int {
	return q.size
}
func main() {
	// Create a circular queue of integers with capacity 5
	queue := NewCircularQueue[int](5)

	for i := 1; i <= 5; i++ {
		if err := queue.Enqueue(i); err != nil {
			fmt.Println("Error:", err)
		}
	}

	// Print front and rear elements
	if front, err := queue.Front(); err == nil {
		fmt.Println("Front element:", front)
	}

	if rear, err := queue.Rear(); err == nil {
		fmt.Println("Rear element:", rear)
	}

	// Dequeue elements
	for i := 1; i <= 3; i++ {
		if value, err := queue.Dequeue(); err == nil {
			fmt.Println("Dequeued:", value)
		}
	}

	// Enqueue more elements
	if err := queue.Enqueue(6); err != nil {
		fmt.Println("Error:", err)
	}
	if err := queue.Enqueue(7); err != nil {
		fmt.Println("Error:", err)
	}

	// Print front and rear elements
	if front, err := queue.Front(); err == nil {
		fmt.Println("Front element:", front)
	}

	if rear, err := queue.Rear(); err == nil {
		fmt.Println("Rear element:", rear)
	}

	// Check size
	fmt.Println("Queue size:", queue.Size())

	// Create a circular queue of strings with capacity 3
	stringQueue := NewCircularQueue[string](3)

	// ing elements
	if err := stringQueue.Enqueue("A"); err != nil {
		fmt.Println("Error:", err)
	}
	if err := stringQueue.Enqueue("B"); err != nil {
		fmt.Println("Error:", err)
	}
	if err := stringQueue.Enqueue("C"); err != nil {
		fmt.Println("Error:", err)
	}

	// Print front and rear elements for string queue
	if front, err := stringQueue.Front(); err == nil {
		fmt.Println("Front element (string):", front)
	}

	if rear, err := stringQueue.Rear(); err == nil {
		fmt.Println("Rear element (string):", rear)
	}
}
