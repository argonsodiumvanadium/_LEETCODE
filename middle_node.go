//https://leetcode.com/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func 
middleNode(head *ListNode) *ListNode {
    fast, slow := head, head

    for i := 1; fast != nil ; i++ {
        if i%2 == 0 {
            slow = slow.Next
        }
        fast = fast.Next
    }

    return slow   
}
