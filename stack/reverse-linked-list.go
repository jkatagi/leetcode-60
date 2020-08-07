/**
* Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
* Memory Usage: 3.6 MB, less than 5.45% of Go online submissions for Reverse Linked List.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	vals := []int{}
	current := head
	// add values
	for {
		vals = append(vals, current.Val)
		if current.Next == nil {
			break
		} else {
			current = current.Next
		}
	}
	reversed := make([]int, 0, len(vals))
	for i := len(vals) - 1; i >= 0; i-- {
		reversed = append(reversed, vals[i])
	}
	return makeListNode(reversed)
}

func makeListNode(vals []int) *ListNode {
	if len(vals) == 1 {
		return &ListNode{
			Val:  vals[0],
			Next: nil,
		}
	} else {
		return &ListNode{
			Val:  vals[0],
			Next: makeListNode(vals[1:]),
		}
	}
}
