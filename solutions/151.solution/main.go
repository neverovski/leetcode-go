package solutions

func reverseWords(s string) string {
	str := []byte(s)
	lastSpaceIndex := len(str)
	out := make([]byte, 0)

	for index := len(str); index >= 0; index-- {
		if index == 0 || str[index-1] == ' ' {
			tmp := str[index:lastSpaceIndex]

			if len(tmp) > 0 {
				out = append(out, tmp...)
				out = append(out, ' ')
			}

			lastSpaceIndex = index - 1
		}
	}

	if out[len(out)-1] == ' ' {
		return string(out[:len(out)-1])
	}

	return string(out)
}
