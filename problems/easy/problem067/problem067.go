package problem067

func addBinary(a string, b string) string {
	i, j, carry, sum := len(a)-1, len(b)-1, 0, ""

	for i >= 0 || j >= 0 {
		temp := carry
		if i >= 0 {
			if a[i] == '1' {
				temp++
			}
			i--
		}

		if j >= 0 {
			if b[j] == '1' {
				temp++
			}
			j--
		}

		carry = temp / 2
		sum = string(byte(temp%2)+'0') + sum
	}

	if carry != 0 {
		sum = "1" + sum
	}

	return sum
}
