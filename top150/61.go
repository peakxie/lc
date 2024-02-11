package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 0
	tail := head
	for node := head; node != nil; node = node.Next {
		tail = node
		n++
	}

	k = n - (k % n)
	tail.Next = head

	for node := tail; ; node = node.Next {
		if k == 0 {
			head = node.Next
			node.Next = nil
			break
		}
		k--
	}

	return head
}
