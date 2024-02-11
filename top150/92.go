package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var l, r *ListNode
	nHead := &ListNode{Next: head}
	for n := nHead; n != nil; n = n.Next {
		if left == 1 {
			l = n
		}
		if right == 1 {
			r = n
		}
		left--
		right--
	}

	rl, _ := reverse(l.Next, r.Next)
	l.Next = rl
	return nHead.Next

}

func reverse(l, r *ListNode) (*ListNode, *ListNode) {
	h, t := l, r
	for l != t {
		lnext := l.Next
		l.Next = r.Next
		r.Next = l
		l = lnext
	}
	return t, h
}
