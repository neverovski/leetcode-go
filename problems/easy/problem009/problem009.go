package problem009

const base = 10

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverseNum := 0
	tempNum := x

	for tempNum > 0 {
		lastDigit := tempNum % base
		reverseNum = reverseNum*base + lastDigit

		tempNum /= base
	}

	return reverseNum == x
}
