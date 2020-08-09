/**
* Runtime: 4 ms, faster than 97.52% of Go online submissions for Path Sum.
* Memory Usage: 4.8 MB, less than 100.00% of Go online submissions for Path Sum.
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumHelper(root, sum)
}

func hasPathSumHelper(root *TreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	}
	sub := sum - root.Val
	var hasSum bool
	if sub == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		hasSum = hasSum || hasPathSumHelper(root.Left, sub)
	}
	if root.Right != nil {
		hasSum = hasSum || hasPathSumHelper(root.Right, sub)
	}
	return hasSum
}
