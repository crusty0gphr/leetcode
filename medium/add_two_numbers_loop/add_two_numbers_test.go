package add_two_numbers_loop

import (
	"fmt"
	"github.com/harout/leetcode/lib"
	"testing"
)

func getMocksOther() (*lib.ListNode, *lib.ListNode) {
	l1 := lib.List{}
	l1.Insert(8)
	l1.Insert(1)

	l2 := lib.List{}
	l2.Insert(0)

	return l1.Head, l2.Head
}

func getMocks() (*lib.ListNode, *lib.ListNode) {
	l1 := lib.List{}
	l1.Insert(0)

	l2 := lib.List{}
	l2.Insert(7)
	l2.Insert(3)

	return l1.Head, l2.Head
}

func getMocksSmall() (*lib.ListNode, *lib.ListNode) {
	l1 := lib.List{}
	l1.Insert(2)
	l1.Insert(4)
	l1.Insert(3)

	l2 := lib.List{}
	l2.Insert(5)
	l2.Insert(6)
	l2.Insert(4)

	return l1.Head, l2.Head
}

func getMocksLarge() (*lib.ListNode, *lib.ListNode) {
	l1 := lib.List{}
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)

	l2 := lib.List{}
	l2.Insert(9)
	l2.Insert(9)
	l2.Insert(9)
	l2.Insert(9)

	return l1.Head, l2.Head
}

func TestAddTwoNumbers_smallSet(t *testing.T) {
	eL := lib.List{}
	eL.Insert(7)
	eL.Insert(0)
	eL.Insert(8)
	e := eL.Head

	l1, l2 := getMocksSmall()
	list := addTwoNumbers(l1, l2)

	for list.Next != nil && e.Next != nil {
		if e.Val != list.Val {
			t.Errorf("f.addTwoNumbers(l1, l2)")
		}
		fmt.Println(list.Val)
		list = list.Next
		e = e.Next
	}
	fmt.Println(list.Val)

	if e.Val != list.Val {
		t.Errorf("f.addTwoNumbers(l1, l2)")
	}
}

func TestAddNumbers_largerSet(t *testing.T) {
	eL := lib.List{}
	eL.Insert(8)
	eL.Insert(9)
	eL.Insert(9)
	eL.Insert(9)
	eL.Insert(0)
	eL.Insert(0)
	eL.Insert(0)
	eL.Insert(1)
	e := eL.Head

	l1, l2 := getMocksLarge()
	list := addTwoNumbers(l1, l2)

	for list.Next != nil && e.Next != nil {
		if e.Val != list.Val {
			t.Errorf("f.addTwoNumbers(l1, l2)")
		}
		fmt.Println(list.Val)
		list = list.Next
		e = e.Next
	}
	fmt.Println(list.Val)
	if e.Val != list.Val {
		t.Errorf("f.addTwoNumbers(l1, l2)")
	}
}

func TestAddNumbers_zero_set(t *testing.T) {
	eL := lib.List{}
	eL.Insert(7)
	eL.Insert(3)
	e := eL.Head

	l1, l2 := getMocks()
	list := addTwoNumbers(l1, l2)

	for list.Next != nil && e.Next != nil {
		if e.Val != list.Val {
			t.Errorf("f.addTwoNumbers(l1, l2)")
		}
		fmt.Println(list.Val)
		list = list.Next
		e = e.Next
	}
	fmt.Println(list.Val)
	if e.Val != list.Val {
		t.Errorf("f.addTwoNumbers(l1, l2)")
	}
}

func TestAddNumbers(t *testing.T) {
	//eL := lib.List{}
	//eL.Insert(7)
	//eL.Insert(3)
	//e := eL.Head

	l1, l2 := getMocksOther()
	list := addTwoNumbers(l1, l2)

	for list.Next != nil {
		//if e.Val != list.Val {
		//	t.Errorf("f.addTwoNumbers(l1, l2)")
		//}
		fmt.Println(list.Val)
		list = list.Next
		//e = e.Next
	}
	fmt.Println(list.Val)
	//if e.Val != list.Val {
	//	t.Errorf("f.addTwoNumbers(l1, l2)")
	//}
}
