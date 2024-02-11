package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	nHead := &ListNode{
		Next: head,
	}
	n := k
	l, r := nHead, nHead
	for node := nHead; node != nil && node.Next != nil; node = node.Next {
		if n == 1 {
			r = node
			nl, nr := reverse(l.Next, r.Next)
			l.Next = nl
			l = nr
			node = nr
			n = k
		}
		n--
	}

	return nHead.Next

}

func reverse(l, r *ListNode) (*ListNode, *ListNode) { // E: undefined: ListNode (typecheck)
	h, t := l, r
	for h != r {
		hn := h.Next
		h.Next = t.Next
		t.Next = h
		h = hn
	}
	return r, l
}
