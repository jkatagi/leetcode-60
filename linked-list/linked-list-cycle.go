/**
* Runtime: 32 ms, faster than 5.43% of Go online submissions for Linked List Cycle.
* Memory Usage: 4.8 MB, less than 50.00% of Go online submissions for Linked List Cycle.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	stock := []int{}
	if head == nil {
		return false
	}
	stock = append(stock, head.Val)
	current := head.Next
	for {
		if current == nil {
			return false
		}
		val := current.Val
		for i := range stock {
			if val == stock[i] && len(stock) > i+1 {
				if current.Next != nil && current.Next.Val == stock[i+1] {
					return true
				}
			}
		}
		stock = append(stock, val)
		current = current.Next
	}
}
