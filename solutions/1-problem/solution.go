package solutions

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}

	for i, num := range nums {
		if val, ok := numsMap[target-num]; ok {
			return []int{val, i}
		}

		numsMap[num] = i
	}

	return []int{}
}
