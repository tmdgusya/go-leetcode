package golangleetcode_test

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const FAILED int = -1

func isBalanced(root *TreeNode) bool {
	return isBalancedSubTree(root) != FAILED
}

func isBalancedSubTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := isBalancedSubTree(node.Left)
	if leftHeight == FAILED {
		return FAILED
	}

	rightHeight := isBalancedSubTree(node.Right)
	if rightHeight == FAILED {
		return FAILED
	}

	if abs(leftHeight-rightHeight) > 1 {
		return FAILED
	}

	return 1 + max(leftHeight, rightHeight)
}

// my first solution
func getMaxHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(getMaxHeight(node.Left), getMaxHeight(node.Right))
}

func abs(result int) int {
	if result < 0 {
		return -result
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestCase_110_1(t *testing.T) {
	tree1 := &TreeNode{Val: 3}
	tree2 := &TreeNode{Val: 9}
	tree3 := &TreeNode{Val: 20}
	tree4 := &TreeNode{Val: 15}
	tree5 := &TreeNode{Val: 17}

	tree1.Left = tree2
	tree1.Right = tree3
	tree3.Left = tree4
	tree3.Right = tree5

	result := isBalanced(tree1)

	if result != true {
		t.Errorf("The Result Must be true")
	}
}

func TestCase_110_2(t *testing.T) {
	tree1 := &TreeNode{Val: 1}
	tree2 := &TreeNode{Val: 2}
	tree3 := &TreeNode{Val: 2}
	tree4 := &TreeNode{Val: 3}
	tree5 := &TreeNode{Val: 3}
	tree6 := &TreeNode{Val: 4}
	tree7 := &TreeNode{Val: 4}

	tree1.Left = tree2
	tree1.Right = tree3

	tree2.Left = tree4
	tree2.Right = tree5

	tree4.Left = tree6
	tree4.Right = tree7

	result := isBalanced(tree1)

	if result != false {
		t.Errorf("The Result Must be false")
	}
}
