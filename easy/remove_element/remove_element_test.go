package remove_element

import (
	"testing"
)

func TestRemoveNums(t *testing.T) {
	e := 1
	nums := []int{1, 1, 2}
	r := removeElement(nums, 1)

	if e != r {
		t.Errorf("f.removeElement(%v) | expects: %v | actual: %v", nums, e, r)
	}
}
