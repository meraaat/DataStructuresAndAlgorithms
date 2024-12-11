package main

import "fmt"

func main() {

	stack := &Stack[string]{}

	stack.Push("Golang")
	stack.Push("DSA")
	stack.Push("Stack")
	stack.Push("Queue")

	fmt.Println("Stack size: ", stack.Size())

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println("Popped item: ", item)
	}

	type Point struct {
		X, Y int
	}

	pointStack := &Stack[Point]{}

	pointStack.Push(Point{X: 1, Y: 2})
	pointStack.Push(Point{X: 3, Y: 4})

	for !pointStack.IsEmpty() {

		point, _ := pointStack.Pop()
		fmt.Println("Popped item: ", point)
	}

	fmt.Println("-----------------------------------------------")

	// Create a stack of integers
	intStack := &StackG[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("Integer Stack size:", intStack.Size())

	topInt, _ := intStack.Peek()
	fmt.Println("Top element (int):", topInt)

	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Println("Popped (int):", val)
	}

	// Create a stack of strings
	stringStack := &StackG[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")

	fmt.Println("\nString Stack size:", stringStack.Size())

	topString, _ := stringStack.Peek()
	fmt.Println("Top element (string):", topString)

	for !stringStack.IsEmpty() {
		val, _ := stringStack.Pop()
		fmt.Println("Popped (string):", val)
	}
}