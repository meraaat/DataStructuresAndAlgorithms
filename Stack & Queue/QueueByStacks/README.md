# Queue using Two Stacks in Go

## Overview

This Go package implements a generic Queue data structure using two stacks, providing an efficient and elegant approach to queue operations with generics.

## Features

- Generic implementation supporting any data type
- Efficient queue operations using two internal stacks
- Methods for:
  - Enqueuing elements
  - Dequeuing elements
  - Peeking at the front element
  - Checking queue size
- Error handling for empty queue scenarios

## Usage Example

```go
// Create a new queue of integers
queue := NewQueue[int]()

// Enqueue elements
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// Get front element without removing
front, _ := queue.Front()
fmt.Println(front)  // Output: 1

// Dequeue elements
value, _ := queue.Dequeue()
fmt.Println(value)  // Output: 1

// Check queue size
size := queue.Size()
fmt.Println(size)   // Output: 2
```

## How It Works

### Two-Stack Queue Mechanism

The implementation uses two internal stacks (`stack1` and `stack2`) to simulate queue behavior:

1. **Enqueue Operation**:
   - Elements are always pushed to `stack1`
   - Time Complexity: O(1)

2. **Dequeue Operation**:
   - If `stack2` is empty, all elements from `stack1` are transferred to `stack2`
   - Elements are reversed during transfer, creating the correct queue order
   - The top element of `stack2` is then removed

### Performance Characteristics

- **Time Complexity**:
  - Enqueue: O(1)
  - Dequeue: Amortized O(1)
  - Front: Amortized O(1)
  - Size: O(1)

- **Space Complexity**: O(n), where n is the number of elements

## Detailed Example

```go

// QueueWithStacks Struct:
// +-------------------+
// | stack1: []T       | → Enqueue elements here
// | stack2: []T       | → Dequeue elements from here
// +-------------------+

// Demonstrate queue operations
queue := NewQueue[int]()

// Enqueue elements
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// Internal state after enqueue:
// stack1: [1, 2, 3]
// stack2: []

// Dequeue first element
first, _ := queue.Dequeue()
// Now:
// stack1: []
// stack2: [3, 2]
```

## Advantages

- Efficient element management
- Generic type support
- Simple implementation
- Minimal overhead

## Limitations

- Not thread-safe
- Slightly more complex than a traditional queue
- Requires occasional element transfer

## Use Cases

- Task scheduling
- Resource allocation management
- Breadth-first search algorithms
- Message processing buffers
