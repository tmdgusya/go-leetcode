package tree_test

import (
	"fmt"
	"testing"
)

type Number interface {
	int | float32 | float64
}

type Node[T Number] struct {
	Val       T
	LeftNode  *Node[T]
	RightNode *Node[T]
}

func (node *Node[T]) Left() *Node[T] {
	return node.LeftNode
}

func (node *Node[T]) Right() *Node[T] {
	return node.RightNode
}

func (selfNode *Node[T]) AddNode(node *Node[T]) {
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

func NewNode[T Number](value T) *Node[T] {
	return &Node[T]{Val: value, LeftNode: nil, RightNode: nil}
}

func (node *Node[T]) Print(depth int) {

	fmt.Println("Self : ", node.Val, " Depth : ", depth, "Addr : ", &node)

	depth++

	if node.LeftNode != nil {
		node.LeftNode.Print(depth)
	}

	if node.RightNode != nil {
		node.RightNode.Print(depth)
	}

}

// Search Target in Tree
// Return Depth of Target
// Time Complexity : O(logN), Bad (O(N))
// Space Complexity: O(logN)
func (node *Node[T]) BinarySearch(target *Node[T], depth int) (d int, isExist bool, targetNode *Node[T]) {
	depth++

	if target.Val == node.Val {
		return depth, true, node
	}

	if target.Val > node.Val {
		return node.RightNode.BinarySearch(target, depth)
	}

	if target.Val < node.Val {
		return node.LeftNode.BinarySearch(target, depth)
	}

	return -1, false, nil
}

func (node *Node[T]) DeleteNode(target *Node[T]) bool {
	_, isExist, targetNode := node.BinarySearch(target, 0)

	if isExist == false || targetNode == nil {
		return false
	}
	successor := targetNode.LeftNode.findSuccessor()

	result := successor.swapNode(targetNode)

	return result
}

func (node *Node[T]) findSuccessor() *Node[T] {
	if node.RightNode == nil {
		return node
	}

	return node.RightNode.findSuccessor()
}

func (node *Node[T]) TraversalPreOrder() {
	if node == nil {
		return
	}

	if node.LeftNode != nil {
		node.LeftNode.TraversalPreOrder()
	}

	fmt.Println(node.Val)

	if node.RightNode != nil {
		node.RightNode.TraversalPreOrder()
	}
}

// If you want deep copy tree, then you use it
// Use it when if you want search root node fisrt, before search leaf nodes
func (node *Node[T]) TraversalInOrder() {
	if node == nil {
		return
	}

	fmt.Println(node.Val)

	if node.LeftNode != nil {
		node.LeftNode.TraversalInOrder()
	}

	if node.RightNode != nil {
		node.RightNode.TraversalInOrder()
	}
}

// first see leef node, then see root node
func (node *Node[T]) TraversalPostOrder() {
	if node == nil {
		return
	}

	if node.LeftNode != nil {
		node.LeftNode.TraversalPostOrder()
	}

	if node.RightNode != nil {
		node.RightNode.TraversalPostOrder()
	}

	fmt.Println(node.Val)
}

func (node *Node[T]) deepCopy() *Node[T] {
	if node == nil {
		return nil
	}

	copyNode := NewNode(node.Val)
	copyNode.LeftNode = node.LeftNode.deepCopy()
	copyNode.RightNode = node.RightNode.deepCopy()

	return copyNode
}

func (node *Node[T]) swapNode(node1 *Node[T]) bool {
	var temp Node[T] // like nil

	node.LeftNode = node1.LeftNode
	node.RightNode = node1.RightNode

	*node1 = *node

	*node = temp
	return true
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

func TestBinarySearch(t *testing.T) {
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

	depth, isExist, _ := node1.BinarySearch(node7, 0)

	if depth != 2 && isExist != true {
		t.Errorf("[Error] 7 is must be exist in Current Tree!")
	}
}

func TestDeleteNode(t *testing.T) {
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

	success := node1.DeleteNode(node6)

	node1.Print(0)

	if success != true && node1.RightNode.Val != 7 {
		t.Errorf("Error")
	}
}

func TestDeleteNode2(t *testing.T) {
	node1 := NewNode(6)
	node2 := NewNode(4)
	node3 := NewNode(10)
	node4 := NewNode(2)
	node5 := NewNode(5)
	node6 := NewNode(1)
	node7 := NewNode(3)
	node8 := NewNode(8)
	node9 := NewNode(7)
	node10 := NewNode(9)
	node11 := NewNode(15)
	node12 := NewNode(12)

	node1.AddNode(node2)
	node1.AddNode(node3)
	node1.AddNode(node4)
	node1.AddNode(node5)
	node1.AddNode(node6)
	node1.AddNode(node7)
	node1.AddNode(node8)
	node1.AddNode(node9)
	node1.AddNode(node10)
	node1.AddNode(node11)
	node1.AddNode(node12)

	success := node1.DeleteNode(node3)
	if success != true && node1.RightNode.Val != 9 {
		t.Errorf("Error")
	}
}

func TestTraversalPreOrder(t *testing.T) {
	node1 := NewNode(6)
	node2 := NewNode(4)
	node3 := NewNode(10)
	node4 := NewNode(2)
	node5 := NewNode(5)
	node6 := NewNode(1)
	node7 := NewNode(3)
	node8 := NewNode(8)
	node9 := NewNode(7)
	node10 := NewNode(9)
	node11 := NewNode(15)
	node12 := NewNode(12)

	node1.AddNode(node2)
	node1.AddNode(node3)
	node1.AddNode(node4)
	node1.AddNode(node5)
	node1.AddNode(node6)
	node1.AddNode(node7)
	node1.AddNode(node8)
	node1.AddNode(node9)
	node1.AddNode(node10)
	node1.AddNode(node11)
	node1.AddNode(node12)

	node1.TraversalPreOrder()

}

func TestTraversalPostOrder(t *testing.T) {
	node1 := NewNode(6)
	node2 := NewNode(4)
	node3 := NewNode(10)
	node4 := NewNode(2)
	node5 := NewNode(5)
	node6 := NewNode(1)
	node7 := NewNode(3)
	node8 := NewNode(8)
	node9 := NewNode(7)
	node10 := NewNode(9)
	node11 := NewNode(15)
	node12 := NewNode(12)

	node1.AddNode(node2)
	node1.AddNode(node3)
	node1.AddNode(node4)
	node1.AddNode(node5)
	node1.AddNode(node6)
	node1.AddNode(node7)
	node1.AddNode(node8)
	node1.AddNode(node9)
	node1.AddNode(node10)
	node1.AddNode(node11)
	node1.AddNode(node12)

	node1.TraversalPostOrder()

}

func TestDeepCopy(t *testing.T) {
	node1 := NewNode(6)
	node2 := NewNode(4)
	node3 := NewNode(10)
	node4 := NewNode(2)
	node5 := NewNode(5)
	node6 := NewNode(1)
	node7 := NewNode(3)
	node8 := NewNode(8)
	node9 := NewNode(7)
	node10 := NewNode(9)
	node11 := NewNode(15)
	node12 := NewNode(12)

	node1.AddNode(node2)
	node1.AddNode(node3)
	node1.AddNode(node4)
	node1.AddNode(node5)
	node1.AddNode(node6)
	node1.AddNode(node7)
	node1.AddNode(node8)
	node1.AddNode(node9)
	node1.AddNode(node10)
	node1.AddNode(node11)
	node1.AddNode(node12)

	copyNode := node1.deepCopy()

	fmt.Print(copyNode)

	if copyNode.LeftNode.LeftNode.Val != node1.LeftNode.LeftNode.Val {
		t.Error("Not Deep Copy...")
	}

}
