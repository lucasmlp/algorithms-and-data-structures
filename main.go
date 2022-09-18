package main

import (
	"log"

	"github.com/machado-br/algorithms-and-data-structures/dataStructures"
)

func main() {
	log.Println("Testing stack")
	testStack()
	log.Println("Testing queue")
	testQueue()
	log.Println("Testing linked list")
	testLinkedList()
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
		log.Printf("element: %v\n", element)
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
		log.Printf("element: %v\n", element)
	}
}

func testLinkedList() {
	linkedList := dataStructures.NewLinkedList("head")
	log.Println("Inserting node_1")
	_ = linkedList.Insert("node_1")

	log.Println("Inserting node_2")
	_ = linkedList.Insert("node_2")

	log.Println("Inserting node_3")
	_ = linkedList.Insert("node_3")

	log.Printf("linkedList: %v\n", linkedList)
	for _, node := range linkedList.List {
		log.Printf("node: %v\n", node)
	}
}
