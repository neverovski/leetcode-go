package problem066

const DIVISOR = 10

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i] + carry
		carry = digit / DIVISOR
		digits[i] = digit % DIVISOR
	}

	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
