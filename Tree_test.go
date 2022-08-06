package golangleetcode

import (
	"fmt"
	"testing"
)

type Node struct {
	Val       int
	LeftNode  *Node
	RightNode *Node
}

func (node *Node) Left() *Node {
	return node.LeftNode
}

func (node *Node) Right() *Node {
	return node.RightNode
}

func (selfNode *Node) AddNode(node *Node) {
	if selfNode.Val > node.Val {
		if selfNode.LeftNode == nil {
			selfNode.LeftNode = node
		}
		if selfNode.LeftNode != nil {
			selfNode.LeftNode.AddNode(node)
		}
	}

	if selfNode.Val < node.Val {
		if selfNode.RightNode == nil {
			selfNode.RightNode = node
		}
		if selfNode.RightNode != nil {
			selfNode.RightNode.AddNode(node)
		}
	}
}

func NewNode(value int) *Node {
	return &Node{Val: value, LeftNode: nil, RightNode: nil}
}

func (node *Node) Print(depth int) {

	fmt.Println("Self : ", node.Val, " Depth : ", depth)

	depth++

	if node.LeftNode != nil {
		fmt.Println("Left Node Value : ", node.LeftNode.Val, " Depth : ", depth)
		node.LeftNode.Print(depth)
	}

	if node.RightNode != nil {
		fmt.Println("Right Node Value : ", node.RightNode.Val, " Depth : ", depth)
		node.RightNode.Print(depth)
	}

}

func TestNode(t *testing.T) {
	node1 := NewNode(5)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(1)
	node6 := NewNode(6)
	node7 := NewNode(7)

	node1.AddNode(node2)
	node1.AddNode(node3)
	node1.AddNode(node4)
	node1.AddNode(node5)
	node1.AddNode(node6)
	node1.AddNode(node7)

	node1.Print(0)
}
