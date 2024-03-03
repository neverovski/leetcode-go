package problem001

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if num, ok := numMap[target-nums[i]]; ok {
			return []int{i, num}
		}

		numMap[nums[i]] = i
	}

	return []int{}
}
