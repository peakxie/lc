package top150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, node *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		carry, sum = sum/10, sum%10
		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			node = head
		} else {
			node.Next = &ListNode{ // E: undefined: ListNode (typecheck)
				Val: sum,
			}
			node = node.Next
		}
	}
	if carry != 0 {
		node.Next = &ListNode{ // E: undefined: ListNode (typecheck)
			Val: carry,
		}
	}
	return head
}
