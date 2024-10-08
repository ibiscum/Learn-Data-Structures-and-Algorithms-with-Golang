// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// Node class
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method of LinkedList
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil {
		//fmt.Println(node.property)
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}

	linkedList.headNode = node

}

// NodeWithValue method of LinkedList
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method of LinkedList
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	// var nodeWith *Node
	var nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		//fmt.Println(node.property)
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}

}

// LastNode method of LinkedList
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method of LinkedList
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	// var lastNode *Node
	var lastNode = linkedList.LastNode()

	if lastNode != nil {

		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

// IterateList method of LinkedList
func (linkedList *LinkedList) IterateList() {

	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {

		fmt.Println(node.property)
	}
}

// NodeBetweenValues method of LinkedList
func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// main method
func main() {

	// var linkedList LinkedList
	var linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)

	// var node *Node
	var node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)

}
