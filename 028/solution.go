package solution_028

import "strings"

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

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

func strStrUsingStrings(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
