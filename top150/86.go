package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	l, r := &ListNode{}, &ListNode{}
	lh, rh := l, r
	for n := head; n != nil; n = n.Next {
		if n.Val < x {
			l.Next = n
			l = l.Next
		} else {
			r.Next = n
			r = r.Next
		}
	}
	r.Next = nil // 这一步不可少，他可能引起循环链表
	l.Next = rh.Next
	return lh.Next
}
