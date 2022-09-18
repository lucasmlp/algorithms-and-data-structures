package main

import (
	"fmt"

	"github.com/machado-br/algorithms-and-data-structures/dataStructures"
)

func main() {
	fmt.Println("Testing stack")
	testStack()
	fmt.Println("Testing queue")
	testQueue()
}

func testStack() {
	stack := dataStructures.NewStack()

	stack.Push("Golang ")
	stack.Push("is ")
	stack.Push("the ")
	stack.Push("best ")
	stack.Push("language!")

	stackLength := len(stack)

	for i := 0; i < stackLength; i++ {
		element, _ := stack.Pop()
		fmt.Printf("element: %v\n", element)
	}
}

func testQueue() {
	queue := dataStructures.NewQueue()

	queue.Enqueue("Golang ")
	queue.Enqueue("is ")
	queue.Enqueue("the ")
	queue.Enqueue("best ")
	queue.Enqueue("language!")

	stackLength := len(queue)

	for i := 0; i < stackLength; i++ {
		element, _ := queue.Dequeue()
		fmt.Printf("element: %v\n", element)
	}
}
