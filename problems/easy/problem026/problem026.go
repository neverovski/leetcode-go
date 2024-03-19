package problem026

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	key := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[key] = nums[i]
			key++
		}
	}

	return key
}
