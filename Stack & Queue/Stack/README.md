### Stack: Time & Space Complexity and Use Cases

#### Time Complexity
| Operation      | Description                       | Time Complexity |
|----------------|-----------------------------------|-----------------|
| **Push**       | Add an element to the top of the stack | O(1) (amortized) |
| **Pop**        | Remove the top element from the stack | O(1)             |
| **Peek**       | View the top element without removal | O(1)             |
| **IsEmpty**    | Check if the stack is empty          | O(1)             |
| **Size**       | Get the number of elements in the stack | O(1)             |

#### Space Complexity
| Aspect                | Description                                     | Space Complexity |
|-----------------------|-------------------------------------------------|------------------|
| **Stack Storage**     | Proportional to the number of elements          | O(n)             |
| **Overhead (Slice)**  | Extra memory for slice capacity during resizing | O(n)             |

#### Use Cases
| Use Case                      | Description                                                                                       |
|-------------------------------|---------------------------------------------------------------------------------------------------|
| **Expression Evaluation**     | Evaluate or convert mathematical expressions (e.g., infix to postfix).                           |
| **Backtracking**              | Undo actions in algorithms like maze solving, Sudoku, or Depth-First Search (DFS).               |
| **Function Call Management**  | Track active function calls in programming languages using a call stack.                         |
| **Undo/Redo Functionality**   | Implement undo/redo features in applications like text editors and graphic design tools.          |
| **Syntax Checking**           | Validate balanced parentheses, brackets, or braces.                                              |
| **Browser Navigation**        | Simulate "back" and "forward" navigation in web browsers using two stacks.                      |
| **String Reversal**           | Reverse strings by pushing characters onto a stack and popping them off.                         |
| **Memory Management**         | Allocate and deallocate memory efficiently in low-level programming.                             |
| **Path Simplification**       | Simplify file system paths (e.g., resolving `"/home/../usr/bin"` to `"/usr/bin"`).              |
| **Game Algorithms**           | Solve problems like the Tower of Hanoi by representing pegs and disks as stacks.                 |

---
Use this chart to understand the **efficiency and versatility** of a stack in various computational scenarios!

