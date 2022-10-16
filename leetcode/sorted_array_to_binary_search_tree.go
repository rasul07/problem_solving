package main

/*
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var result *TreeNode

	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	i := 0
	j := len(nums) - 1
	mid := (i + j) / 2

	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[mid+1:])

	result = &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}

	return result
}

// func main() {
// 	nums := []int{-10, -3, 0, 5, 9}

// 	fmt.Println(sortedArrayToBST(nums))
// }
