/**
* Runtime: 4 ms, faster than 97.60% of Go online submissions for Minimum Depth of Binary Tree.
* Memory Usage: 5.3 MB, less than 76.92% of Go online submissions for Minimum Depth of Binary Tree.
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	return findMinDepth(root)
}

func findMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + findMinDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + findMinDepth(root.Left)
	}
	return 1 + min(findMinDepth(root.Left), findMinDepth(root.Right))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
