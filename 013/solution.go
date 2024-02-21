package solution_013

var symbols = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func RomanToInt(s string) int {
	num := 0

	for i := 0; i < len(s); i++ {
		if i > 0 && symbols[s[i]] > symbols[s[i-1]] {
			num += symbols[s[i]] - 2*symbols[s[i-1]]
		} else {
			num += symbols[s[i]]
		}
	}

	return num
}
