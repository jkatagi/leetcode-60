/**
 * Runtime: 108 ms, faster than 76.19% of Go online submissions for Convert Sorted Array to Binary Search Tree.
 * Memory Usage: 249.8 MB, less than 5.02% of Go online submissions for Convert Sorted Array to Binary Search Tree.
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	middle := (start + end) / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBSTHelper(nums, start, middle-1),
		Right: sortedArrayToBSTHelper(nums, middle+1, end),
	}
} 
