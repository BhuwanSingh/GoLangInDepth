package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type SinglyLinkedList struct {
	headnode *Node
}

// method to add an element to the beginning of the linked list.
func (singlyLinkedList *SinglyLinkedList) AddToTheRoot(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if singlyLinkedList.headnode != nil {
		node.nextNode = singlyLinkedList.headnode
	}

	singlyLinkedList.headnode = node
}

func (singlyLinkedList *SinglyLinkedList) FindNode(property int) *Node {
	var node *Node
	var nodeWith *Node

	for node = singlyLinkedList.headnode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (singlyLinkedList *SinglyLinkedList) AddAfterThisElement(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	nodeWith := singlyLinkedList.FindNode(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func (singlyLinkedList *SinglyLinkedList) GetLastNode() *Node {
	for node := singlyLinkedList.headnode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (singlyLinkedList *SinglyLinkedList) AddToLastNode(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	lastnode := singlyLinkedList.GetLastNode()

	if lastnode != nil {
		lastnode.nextNode = node
	}
}

func (singlyLinkedList *SinglyLinkedList) ListIterator() {
	for node := singlyLinkedList.headnode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	var singlyLinkedList SinglyLinkedList
	singlyLinkedList = SinglyLinkedList{}
	singlyLinkedList.AddToTheRoot(1)
	singlyLinkedList.AddToTheRoot(2)
	singlyLinkedList.AddToLastNode(4)
	singlyLinkedList.AddAfterThisElement(1, 6)
	singlyLinkedList.ListIterator()
}
