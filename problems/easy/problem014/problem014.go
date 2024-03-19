package problem014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	if len(strs) == 1 {
		return strs[0]
	}

	return findCommonPrefix(strs)
}

func findCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
