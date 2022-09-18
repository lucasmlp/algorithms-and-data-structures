package dataStructures

import (
	"log"
)

type LinkedList struct {
	List []Node
	Head *Node
}

type Node struct {
	Key  string
	Next *Node
}

func NewLinkedList(key string) LinkedList {
	head := Node{
		Key: key,
	}

	return LinkedList{
		Head: &head,
		List: []Node{head},
	}
}

func (lL *LinkedList) Insert(key string) *Node {
	node := lL.Head
	for node.Next != nil {
		node = node.Next
	}

	newNode := Node{
		Key: key,
	}
	log.Printf("newNode: %v\n", newNode)

	node.Next = &newNode
	log.Printf("node: %v\n", node)

	lL.List = append(lL.List, newNode)

	return &newNode
}
