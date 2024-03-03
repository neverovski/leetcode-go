package problem066

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i] + carry
		carry = digit / 10
		digits[i] = digit % 10
	}

	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
