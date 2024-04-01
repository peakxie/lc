package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func merge(h1, h2 *ListNode) *ListNode {
	head := &ListNode{}
	n := head
	for h1 != nil && h2 != nil {
		if h1.Val > h2.Val {
			n.Next = h2
			h2 = h2.Next
		} else {
			n.Next = h1
			h1 = h1.Next
		}
		n = n.Next
	}
	if h1 != nil {
		n.Next = h1
	} else {
		n.Next = h2
	}
	return head.Next
}

func sort(h, t *ListNode) *ListNode {
	if h == t {
		return nil
	}

	if h.Next == t {
		h.Next = nil
		return h
	}

	fast, slow := h, h
	for fast != t {
		slow = slow.Next
		fast = fast.Next
		if fast != t {
			fast = fast.Next
		}
	}

	return merge(sort(h, slow), sort(slow, t))
}
