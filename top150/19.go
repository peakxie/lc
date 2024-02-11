package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	stack := []*ListNode{}
	h := &ListNode{
		Next: head,
	}
	for node := h; node != nil; node = node.Next {
		stack = append(stack, node)
	}

	node := stack[len(stack)-n-1]
	node.Next = node.Next.Next

	return h.Next
}
