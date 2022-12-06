// https://leetcode.com/problems/reverse-linked-list/
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head

	for i := 0; fast != nil; i++ {
		if i%2 == 1 {
			slow = (*slow).Next
		}

		fast = (*fast).Next
	}

	reversed := reverseList(slow)

	for i := 0; reversed != nil; i++ {
		if (*head).Val != (*reversed).Val {
			return false
		}
		reversed = reversed.Next
		head = head.Next
	}

	return true
}

func reverseList(itr *ListNode) (reverse *ListNode) {
	for itr != nil {
		temp := itr
		itr = (*itr).Next
		(*temp).Next = reverse
		reverse = temp
	}

	return
}