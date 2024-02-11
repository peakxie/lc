package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	h := &ListNode{
		Next: head,
	}

	for n := h; n.Next != nil && n.Next.Next != nil; {
		if n.Next.Val == n.Next.Next.Val {
			m := n.Next.Next
			for m.Next != nil && m.Val == m.Next.Val {
				m = m.Next
			}
			n.Next = m.Next
		} else {
			n = n.Next
		}
	}
	return h.Next
}
