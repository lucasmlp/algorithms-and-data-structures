package dataStructures

type LinkedList struct {
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

	node.Next = &newNode

	return &newNode
}

func (lL *LinkedList) Delete(key string) {
	node := lL.Head
	var lastNode *Node
	for node.Key != key{
		lastNode = node
		node = node.Next
	}

	lastNode.Next = node.Next

	node = nil
}

func (lL *LinkedList) Search(key string) *Node {
	node := lL.Head
	for node.Key != key{
		node = node.Next
	}
	
	return node
}