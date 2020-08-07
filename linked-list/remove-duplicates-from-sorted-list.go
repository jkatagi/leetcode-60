/ **
* Runtime: 4 ms, faster than 89.31% of Go online submissions for Remove Duplicates from Sorted List.
* Memory Usage: 3.5 MB, less than 10.00% of Go online submissions for Remove Duplicates from Sorted List.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    uniqVals := []int{}
    uniqVals = append(uniqVals, head.Val)
    current := head
    for {
        if current.Next == nil {
            if !isAlreadyAppeared(current.Val, uniqVals) {
                uniqVals = append(uniqVals, current.Val)
            }
            break
        }
        if isAlreadyAppeared(current.Val, uniqVals) {
            current = current.Next
        } else {
            uniqVals = append(uniqVals, current.Val)
        }
    }
    return makeListNode(uniqVals)
}

func isAlreadyAppeared(val int, stock []int) bool{
    for i := range stock {
        if val == stock[i] {
            return true
        }
    }
    return false
}

func makeListNode(vals []int) *ListNode {
    if len(vals) == 1 {
        return &ListNode{
            Val: vals[0],
            Next: nil,
        }
    } else {
        return &ListNode {
            Val: vals[0],
            Next: makeListNode(vals[1:]),
        }
    }
}
