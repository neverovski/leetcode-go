package solutions

func reverseWords(s string) string {
	str := []byte(s)
	length := len(str)
	lastSpaceIndex := -1

	for index := 0; index <= length; index++ {
		if index == length || str[index] == ' ' {
			startIndex := lastSpaceIndex + 1
			endIndex := index - 1

			for startIndex < endIndex {
				tmp := str[startIndex]
				str[startIndex] = str[endIndex]
				str[endIndex] = tmp

				startIndex++
				endIndex--
			}

			lastSpaceIndex = index
		}
	}

	return string(str)
}
