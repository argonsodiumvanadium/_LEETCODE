// https://leetcode.com/problems/reverse-linked-list/
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

/*  creating new list */
func 
reverseList(itr *ListNode) (reverse *ListNode) {
	for ;itr != nil; itr=itr.Next {
		reverse = &ListNode{(*itr).Val ,reverse}
	}
	
	return
}


/* inplace reversal */
func 
reverseList(itr *ListNode) (reverse *ListNode) {
	for itr != nil {
		temp:=itr
		itr=(*itr).Next
		(*temp).Next=reverse 
		reverse=temp
	}
	
	return
}

