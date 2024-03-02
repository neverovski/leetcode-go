package solution_58

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	lenLastWord := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(' ') {
			break
		}

		lenLastWord++
	}

	return lenLastWord
}
