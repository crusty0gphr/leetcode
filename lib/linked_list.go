package lib

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	Len  int
}

func (list *List) Insert(v int) {
	node := ListNode{}
	node.Val = v

	if list.Len == 0 {
		list.Head = &node
		list.Len++
		return
	}

	ptr := list.Head

	for i := 0; i < list.Len; i++ {
		if ptr.Next == nil {
			ptr.Next = &node
			list.Len++
			return
		}
		ptr = ptr.Next
	}
}

func (list *List) InsertAt(pos int, v int) {
	if pos < 0 || pos > list.Len {
		return
	}

	node := ListNode{}
	node.Val = v

	if pos == 0 {
		prevHead := list.Head
		node.Next = prevHead
		list.Head = &node
		list.Len++
		return
	}

	node.Next = list.Get(pos)
	prevNode := list.Get(pos - 1)
	prevNode.Next = &node
	list.Len++
}

func (list *List) Delete(v int) error {
	if list.Len == 0 {
		fmt.Println("List is empty")
		return errors.New("list is empty")
	}

	ptr := list.Head
	for i := 0; i < list.Len; i++ {
		if ptr.Val == v {
			if i == 0 {
				list.Head = ptr.Next
				return nil
			}

			ptr = list.Get(i)
			prevPtr := list.Get(i - 1)
			prevPtr.Next = ptr.Next
			return nil
		}
		ptr = ptr.Next
	}
	fmt.Println("Not found")
	return nil
}

func (list *List) Get(pos int) *ListNode {
	if pos < 0 {
		return nil
	}

	if pos == 0 {
		return list.Head
	}

	ptr := list.Head
	for i := 0; i < pos; i++ {
		ptr = ptr.Next
	}
	return ptr
}

func (list *List) Print() {
	if list.Len == 0 {
		fmt.Println("empty")
	}

	ptr := list.Head
	for i := 0; i < list.Len; i++ {
		fmt.Println("*ListNode: ", ptr.Val)
		ptr = ptr.Next
	}
}
