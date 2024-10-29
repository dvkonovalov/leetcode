package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	mp := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := mp[head]; ok {
			return true
		}
		mp[head] = true
		head = head.Next
	}
	return false
}
