package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicatesBasic(t *testing.T) {
	nums := []int{1, 1, 2}
	expectedNums := []int{1, 2}

	k := removeDuplicates(nums)

	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			t.Errorf("f.removeDuplicates(%v) | expects: %v == %v", nums, nums[i] , expectedNums[i])
		}
	}
}

func TestRemoveDuplicatesLarge(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 4, 5, 5, 5, 6, 7, 8, 8, 8, 8}
	//expectedNums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	_ = removeDuplicatesRecursion(nums)

	//for i := 0; i < k; i++ {
	//	if nums[i] != expectedNums[i] {
	//		t.Errorf("f.removeDuplicates(%v) | expects: %v == %v", nums, nums[i] , expectedNums[i])
	//	}
	//}
}
