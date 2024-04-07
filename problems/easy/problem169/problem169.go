package problem169

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hash := make(map[int]int, 0)
	element := nums[0]

	for _, num := range nums {
		hash[num]++

		if hash[num] > hash[element] {
			element = num
		}
	}

	return element
}
