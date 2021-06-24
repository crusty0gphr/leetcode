package remove_element

import "testing"

func TestRemoveNums(t *testing.T) {
	nums := []int{1, 1, 2}
	_ = removeElement(nums, 1)
}
