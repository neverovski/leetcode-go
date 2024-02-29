package solution_026

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	key := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[key] = nums[i]
			key += 1
		}
	}

	return key
}
