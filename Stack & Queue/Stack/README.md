# Stack Implementation in Go: Time & Space Complexity and Use Cases

## Time Complexity

| Operation      | Description                                         | Time Complexity |
|----------------|-----------------------------------------------------|-----------------|
| **Push**       | Add an element to the top of the stack              | O(1) (amortized) |
| **Pop**        | Remove the top element from the stack               | O(1)             |
| **Peek**       | View the top element without removal                | O(1)             |
| **IsEmpty**    | Check if the stack is empty                         | O(1)             |
| **Size**       | Get the number of elements in the stack             | O(1)             |

## Space Complexity

| Aspect                | Description                                       | Space Complexity |
|-----------------------|---------------------------------------------------|------------------|
| **Stack Storage**      | Proportional to the number of elements            | O(n)             |
| **Overhead**           | Extra memory for overhead (pointer or slice)      | O(n)             |

## Comparison Between Array-Based and Linked List-Based Stack Implementations

### Array-Based Stack

**Description:**
An array-based stack uses a dynamically resizable array to store elements. When elements are added, they are appended to the end of the array. If the array reaches its capacity, a new array with double the size is allocated, and the elements are copied over. This resizing is done in an amortized manner, making the average time complexity of **Push** O(1).

**Time Complexity:**
- `Push`: **O(1)** (amortized). When resizing occurs, the complexity becomes O(n), but this happens infrequently.
- `Pop`: **O(1)**. Removes the last element, which is always at the top.
- `Peek`: **O(1)**. Retrieves the top element without modifying the stack.
- `IsEmpty`: **O(1)**. Checks if the stack has any elements.
- `Size`: **O(1)**. Returns the number of elements in the stack.

**Space Complexity:**
- **Stack Storage**: O(n), where `n` is the number of elements in the stack.
- **Overhead**: The extra space required to store elements in the array, and if the array needs resizing, the space complexity temporarily increases.

**Advantages:**
- Better cache locality due to contiguous memory allocation.
- Constant time complexity for most operations (amortized).

**Disadvantages:**
- May waste memory due to resizing when the allocated array is larger than necessary.
- Resizing can have a costly O(n) time complexity during the growth of the array.
- Fixed array size at the time of initialization may limit flexibility.

### Linked List-Based Stack

**Description:**
A linked list-based stack uses a linked list where each node contains an element and a reference (pointer) to the next node in the stack. The stack operates on the head of the list, where **Push** adds an element at the head and **Pop** removes the head node.

**Time Complexity:**
- `Push`: **O(1)**. A new node is added to the head of the linked list.
- `Pop`: **O(1)**. The head node is removed and the next node becomes the new head.
- `Peek`: **O(1)**. Retrieves the value at the head without modifying the stack.
- `IsEmpty`: **O(1)**. Checks if the stack has any nodes.
- `Size`: **O(1)**. Keeps track of the number of nodes.

**Space Complexity:**
- **Stack Storage**: O(n), where `n` is the number of elements in the stack.
- **Overhead**: Each node requires extra memory for the `next` pointer/reference in addition to storing the data.

**Advantages:**
- Dynamic sizing: No need to preallocate memory or resize, as nodes are created as needed.
- No wasted memory compared to array-based stacks with fixed capacities.

**Disadvantages:**
- Extra memory overhead per node due to the `next` pointer.
- Worse cache locality compared to array-based stacks because nodes are stored in non-contiguous memory locations.
- Linked list traversal for other operations may be less efficient in some contexts.

### Efficiency Comparison

| **Aspect**              | **Array-Based Stack**                          | **Linked List-Based Stack**                      |
|-------------------------|------------------------------------------------|--------------------------------------------------|
| **Time Complexity**      | `Push`: O(1) (amortized), `Pop`: O(1), `Peek`: O(1), `IsEmpty`: O(1), `Size`: O(1) | `Push`: O(1), `Pop`: O(1), `Peek`: O(1), `IsEmpty`: O(1), `Size`: O(1) |
| **Space Complexity**     | O(n) (with potential resizing overhead)       | O(n) (with extra overhead for pointers)          |
| **Memory Management**    | Fixed-size array with resizing overhead       | Dynamic memory allocation per node               |
| **Cache Efficiency**     | Better cache locality                         | Worse cache locality due to scattered nodes      |
| **Overhead**             | No extra memory overhead (except resizing)    | Extra memory for storing pointers in each node  |
| **Resizing Cost**        | Potential **O(n)** when resizing (amortized O(1)) | No resizing cost (always O(1))                   |

## Use Cases

| Use Case                      | Description                                                                                       |
|-------------------------------|---------------------------------------------------------------------------------------------------|
| **Expression Evaluation**     | Evaluate or convert mathematical expressions (e.g., infix to postfix).                           |
| **Backtracking**              | Undo actions in algorithms like maze solving, Sudoku, or Depth-First Search (DFS).               |
| **Function Call Management**  | Track active function calls in programming languages using a call stack.                         |
| **Undo/Redo Functionality**   | Implement undo/redo features in applications like text editors and graphic design tools.          |
| **Syntax Checking**           | Validate balanced parentheses, brackets, or braces.                                              |
| **Browser Navigation**        | Simulate "back" and "forward" navigation in web browsers using two stacks.                       |
| **String Reversal**           | Reverse strings by pushing characters onto a stack and popping them off.                         |
| **Memory Management**         | Allocate and deallocate memory efficiently in low-level programming.                             |
| **Path Simplification**       | Simplify file system paths (e.g., resolving `"/home/../usr/bin"` to `"/usr/bin"`).               |
| **Game Algorithms**           | Solve problems like the Tower of Hanoi by representing pegs and disks as stacks.                 |

## Conclusion

This chart provides an overview of the **efficiency and versatility** of a stack in various computational scenarios. Both array-based and linked list-based stack implementations have their strengths and weaknesses. The choice between them depends on the specific use case and performance requirements.

