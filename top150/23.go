package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	return merge(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func merge(h1, h2 *ListNode) *ListNode {

	h := &ListNode{}
	n := h

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

	return h.Next
}
