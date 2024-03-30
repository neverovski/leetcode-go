package problem125

import (
	"strings"
)

const maxDigitChar = '9'
const minLowercaseChar = 'a'
const maxLowercaseChar = 'z'

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; {
		s1, s2 := s[i], s[j]

		if !isAlphanumeric(s1) {
			i++
			continue
		}

		if !isAlphanumeric(s2) {
			j--
			continue
		}

		if s1 != s2 {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= '0' && b <= maxDigitChar) || (b >= minLowercaseChar && b <= maxLowercaseChar)
}
