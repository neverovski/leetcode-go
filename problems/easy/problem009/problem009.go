package problem009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverseNum := 0
	tempNum := x

	for tempNum > 0 {
		lastDigit := tempNum % 10
		reverseNum = reverseNum*10 + lastDigit

		tempNum /= 10
	}

	return reverseNum == x
}
