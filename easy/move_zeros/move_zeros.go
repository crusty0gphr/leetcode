package move_zeros

func moveZeroes(nums []int) {
	var j int
	var zeroes int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		} else {
			zeroes++
		}
	}

	for i := 0; i < zeroes; i++ {
		last := len(nums) - 1
		nums[last-i] = 0
	}
}
