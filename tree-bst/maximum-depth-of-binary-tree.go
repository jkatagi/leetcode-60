/**
* Runtime: 4 ms, faster than 94.23% of Go online submissions for Maximum Depth of Binary Tree.
* Memory Usage: 4.4 MB, less than 85.02% of Go online submissions for Maximum Depth of Binary Tree.
 */
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
*}
 */
func maxDepth(root *TreeNode) int {
	return findMaxDepth(root)

}

func findMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0

	}
	leftDepth := findMaxDepth(root.Left)
	rightDepth := findMaxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1

	}
	return rightDepth + 1

}
