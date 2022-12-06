//https://leetcode.com/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/* implementation 2 */
func removeElements(head *ListNode, val int) *ListNode {
	itr := &ListNode{val + 1, head}
	head = itr

	for itr != nil {
		if itr.Next != nil && itr.Next.Val == val {
			itr.Next = itr.Next.Next
		} else {
			itr = itr.Next
		}
	}

	return head.Next
}

/* implementation 1 */
func removeElements(head *ListNode, val int) *ListNode {
	var previous_itr *ListNode

	itr := head
	previous_itr = nil

	for itr != nil {
		if itr.Val == val {
			if previous_itr == nil {
				head = head.Next
				continue
			} else {
				(*previous_itr).Next = itr.Next
				itr = itr.Next
				continue
			}
		}

		itr = (*itr).Next
		previous_itr = itr
	}

	return head
}

