/**
* Runtime: 28 ms, faster than 90.75% of Go online submissions for Merge Two Binary Trees.
* Memory Usage: 8.6 MB, less than 60.35% of Go online submissions for Merge Two Binary Trees.
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return mergeTreesHelper(t1, t2)
}

func mergeTreesHelper(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 != nil {
		return &TreeNode{
			Val:   t2.Val,
			Left:  t2.Left,
			Right: t2.Right,
		}
	} else if t1 != nil && t2 == nil {
		return &TreeNode{
			Val:   t1.Val,
			Left:  t1.Left,
			Right: t1.Right,
		}
	} else if t1 == nil && t2 == nil {
		return nil
	}
	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTreesHelper(t1.Left, t2.Left),
		Right: mergeTreesHelper(t1.Right, t2.Right),
	}
}

