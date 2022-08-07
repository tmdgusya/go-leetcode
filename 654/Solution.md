# My Solution

![image](https://user-images.githubusercontent.com/57784077/183288872-59a76f49-5d4c-43cb-b747-ea5b4bd984fc.png)

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