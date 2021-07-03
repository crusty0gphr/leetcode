package add_two_numbers

import (
	"github.com/harout/leetcode/lib"
)

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	var c int
	l1.Val += l2.Val
	c = l1.Val / 10
	l1.Val = l1.Val % 10

	if l1.Next == nil && l2.Next != nil {
		l1.Next = &lib.ListNode{Val: 0}
	}

	if c != 0 && l1.Next == nil {
		l1.Next = &lib.ListNode{Val: c}
		return l1
	}

	if l1.Next != nil {
		l1.Next.Val += c

		if l2.Next == nil {
			l2.Next = &lib.ListNode{Val: 0}
		}
		addTwoNumbers(l1.Next, l2.Next)
	}

	return l1
}
