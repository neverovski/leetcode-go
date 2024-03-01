package solution_028

func StrStr(haystack string, needle string) int {
	startIndex := -1

	for i := 0; i < len(haystack); i++ {
		endIndex := i + len(needle)

		if len(haystack) > endIndex-1 && haystack[i:endIndex] == needle {
			startIndex = i

			break
		}
	}

	return startIndex
}

// end = 0 + 1 - 1
// 1 > 1 - false
