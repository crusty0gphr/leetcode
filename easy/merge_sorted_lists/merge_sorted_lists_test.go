package merge_sorted_lists

import (
	"github.com/harout/leetcode/lib"
	"testing"
)

func TestMergeSortedLists(t *testing.T) {
	l1 := lib.List{}
	l1.Insert(4)
	l1.Insert(9)
	l1.Insert(12)

	l2 := lib.List{}
	l2.Insert(1)
	l2.Insert(4)
	l2.Insert(33)

	e1 := lib.List{}
	e1.Insert(1)
	e1.Insert(4)
	e1.Insert(4)
	e1.Insert(9)
	e1.Insert(12)
	e1.Insert(33)

	an1 := mergeTwoLists(l1.Get(0), l2.Get(0))
	al1 := lib.List{Head: an1}

	if !compareLists(e1, al1) {
		t.Errorf("f.mergeTwoLists(l1, l2) | expects: %v | actual: %v", e1, an1)
	}

	e2 := lib.List{}
	an2 := mergeTwoLists(l1.Get(0), l2.Get(0))
	al2 := lib.List{Head: an1}

	if !compareLists(e2, al2) {
		t.Errorf("f.mergeTwoLists(l1, l2) | expects: %v | actual: %v", e2, an2)
	}
}

func compareLists(e lib.List, a lib.List) bool {
	eHead := e.Head
	aHead := a.Head

	for eHead != nil {
		if eHead.Val != aHead.Val {
			return false
		}
		eHead = eHead.Next
		aHead = aHead.Next
	}

	return true
}
