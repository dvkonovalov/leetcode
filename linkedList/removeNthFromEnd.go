package linkedList

func lenList(node *ListNode) int {
	if node == nil {
		return 0
	}
	return lenList(node.Next) + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	counter := lenList(head)
	if counter < 2 {
		return nil
	}
	counter = counter - n - 1
	if counter < 0 {
		return head.Next
	}

	tree := head
	for counter > 0 {
		tree = tree.Next
		counter--
	}
	tree.Next = tree.Next.Next
	return head
}
