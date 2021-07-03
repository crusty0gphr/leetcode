package merge_sorted_lists

import "github.com/harout/leetcode/lib"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	node := new(lib.ListNode)

	if l1.Val <= l2.Val {
		node.Val = l1.Val
		node.Next = mergeTwoLists(l1.Next, l2)
	} else {
		node.Val = l2.Val
		node.Next = mergeTwoLists(l1, l2.Next)
	}

	return node
}
