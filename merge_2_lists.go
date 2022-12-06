//https://leetcode.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(itr1 *ListNode, itr2 *ListNode) *ListNode {
	if itr1 == nil {
		return itr2
	} else if itr2 == nil {
		return itr1
	}

	var head *ListNode
	var tail *ListNode
	if itr1.Val > itr2.Val {
		head = itr2
		tail = itr2
	} else {
		head = itr1
		tail = itr1
	}

	for ; tail != nil; tail = tail.Next {
		if itr1 == nil {
			tail.Next = itr2
			return head
		} else if itr2 == nil {
			tail.Next = itr1
			return head
		} else if itr1.Val > itr2.Val {
			prev := itr2
			itr2 = itr2.Next
			tail.Next = prev
		} else {
			prev := itr1
			itr1 = itr1.Next
			tail.Next = prev
		}
	}

	return head
}