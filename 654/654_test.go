package solution_654

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	pivot := getMaxValudeIndex(nums)
	root := &TreeNode{Val: nums[pivot]}
	process(root, nums, nums[:pivot], nums[pivot+1:])

	return root
}

func process(root *TreeNode, nums []int, left []int, right []int) {

	leftNodeIndex := getMaxValudeIndex(left)
	rightNodeIndex := getMaxValudeIndex(right)

	if leftNodeIndex == -1 {
		root.Left = nil
	} else {
		root.Left = &TreeNode{Val: left[leftNodeIndex]}
	}

	if rightNodeIndex == -1 {
		root.Right = nil
	} else {
		root.Right = &TreeNode{Val: right[rightNodeIndex]}
	}

	if len(left) > 1 {
		process(root.Left, nums, left[:leftNodeIndex], left[leftNodeIndex+1:])
	}

	if len(right) > 1 {
		process(root.Right, nums, right[0:rightNodeIndex], right[rightNodeIndex+1:])
	}
}

func getMaxValudeIndex(nums []int) int {
	maxValue := -1
	maxValueIndex := -1

	for i, v := range nums {
		if max(maxValue, v) != maxValue {
			maxValueIndex = i
			maxValue = v
		}
	}

	return maxValueIndex
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test654_1(t *testing.T) {

	result := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})

	if result.Val != 6 {
		t.Errorf("Error")
	}

	fmt.Println(result)

	firstLeft := result.Left
	firstRight := result.Right
	fmt.Println("[ROOT-LEFT-1] : ", firstLeft)
	fmt.Println("[ROOT-RIGHT-1] : ", firstRight)

	secondLeft := firstLeft.Left
	secondRight := firstLeft.Right
	fmt.Println("[ROOTLEFT-LEFT-1] : ", secondLeft)
	fmt.Println("[ROOTLEFT-RIGHT-1] : ", secondRight)

	secondRightLeft := firstRight.Left
	secondRightRight := firstRight.Right
	fmt.Println("[ROOTRIGHT-LEFT-1] : ", secondRightLeft)
	fmt.Println("[ROOTRIGHT-RIGHT-1] : ", secondRightRight)

	thirdLeft := secondRight.Left
	thirdRight := secondRight.Right
	fmt.Println("[ROOTLEFT-LEFT-2] : ", thirdLeft)
	fmt.Println("[ROOTLEFT-RIGHT-2] : ", thirdRight)
}
