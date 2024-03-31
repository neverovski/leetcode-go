package problem136

func singleNumber(nums []int) int {
	num := make(map[int]struct{}, len(nums)/2)

	for i := 0; i < len(nums); i++ {
		_, ok := num[nums[i]]

		if ok {
			delete(num, nums[i])
		} else {
			num[nums[i]] = struct{}{}
		}
	}

	for k := range num {
		return k
	}

	return 0
}

func singleNumberUsingBit(nums []int) int {
	num := 0

	for i := 0; i < len(nums); i++ {
		num = num ^ nums[i]
	}

	return num
}
