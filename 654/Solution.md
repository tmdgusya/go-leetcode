# My Solution

![image](https://user-images.githubusercontent.com/57784077/183288872-59a76f49-5d4c-43cb-b747-ea5b4bd984fc.png)

```go
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

```

# Best Solution

```go
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0{
        return nil
    }
    var root TreeNode
    max := nums[0]
    num := 0
    for i, v := range nums{
        if max < v{
            max = v
            num = i
        }
    }
    root.Val = max
    root.Left = constructMaximumBinaryTree(nums[:num])
    root.Right = constructMaximumBinaryTree(nums[num + 1:])
    return &root
}``
```

- 내 풀이를 되게 간소하게 만든 버전...