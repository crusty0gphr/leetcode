package remove_element

import "fmt"

func removeElement(nums []int, val int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	fmt.Println(nums)
	return j
}
