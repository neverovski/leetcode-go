package problem168

func convertToTitle(columnNumber int) string {
	title := ""

	if columnNumber < 1 {
		return title
	}

	maxNumber := 26

	for columnNumber > 0 {
		columnNumber--

		title = string(byte('A'+columnNumber%maxNumber)) + title

		columnNumber /= maxNumber
	}

	return title
}
