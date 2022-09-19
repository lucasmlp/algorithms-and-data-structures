package main

import (
	"log"
	"os"

	"github.com/machado-br/algorithms-and-data-structures/dataStructures"
)

func main() {
	action := os.Args[1]

	switch action {
	case "TestStack":
	testStack()

	case "TestQueue":
	testQueue()

	case "TestLinkedList":
	testLinkedList()
}
}

func testStack() {
	log.Println("Testing stack")

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
	log.Println("Testing queue")

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
	log.Println("Testing linked list")

	linkedList := dataStructures.NewLinkedList("head")

	node1 := linkedList.Insert("node_1")
	node2 := linkedList.Insert("node_2")
	_ = linkedList.Insert("node_3")

	linkedList.Delete(node2.Key)

	node := linkedList.Head
	log.Printf("node: %v\n", node)

	for node.Next != nil {
		node = node.Next
		log.Printf("node: %v\n", node)
	}

	nodeFound := linkedList.Search(node1.Key)
	log.Printf("nodeFound: %v\n", nodeFound)
}
