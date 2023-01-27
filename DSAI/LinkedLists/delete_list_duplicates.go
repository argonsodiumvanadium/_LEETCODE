//https://leetcode.com/problems/remove-duplicates-from-sorted-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	for itr := head; itr != nil; {
		if itr.Next != nil && itr.Val == itr.Next.Val {
			itr.Next = itr.Next.Next
		} else {
			itr = itr.Next
		}
	}

	return head
}