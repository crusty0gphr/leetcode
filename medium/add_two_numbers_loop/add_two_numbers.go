package add_two_numbers_loop

import "github.com/harout/leetcode/lib"

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	c, h := 0, l1
	for {
		l1.Val += l2.Val + c
		c = l1.Val / 10
		l1.Val = l1.Val % 10

		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}

		l1, l2 = l1.Next, l2.Next
	}

	for c != 0 {
		if l1.Next == nil {
			l1.Next = &lib.ListNode{Val: 0, Next: nil}
		}
		l1.Next.Val += c
		c = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10
		l1 = l1.Next
	}
	return h
}
