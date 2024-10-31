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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := 0
	ans := ListNode{}
	tmp := &ans
	for l1 != nil && l2 != nil {
		tmp.Next = &ListNode{}
		tmp = tmp.Next

		if l1.Val+l2.Val+prev > 9 {
			tmp.Val = l1.Val + l2.Val + prev - 10
			prev = 1
		} else {
			tmp.Val = l1.Val + l2.Val + prev
			prev = 0
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}

	if prev > 0 {
		if tmp.Next == nil {
			tmp.Next = &ListNode{}
			tmp.Next.Val = 1
			return ans.Next
		}
		tmp = tmp.Next

		for prev > 0 {
			if tmp.Val+1 > 9 {
				tmp.Val = 0
			} else {
				tmp.Val = tmp.Val + 1
				prev = 0
			}
			if tmp.Next == nil && prev > 0 {
				tmp.Next = &ListNode{}
				tmp.Next.Val = 1
				prev = 0
			}
			tmp = tmp.Next
		}
	}

	return ans.Next
}
