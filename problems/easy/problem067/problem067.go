package problem067

const divisor = 2

type BinaryStrings struct {
	a, b string
	i, j *int
}

func addBinary(a string, b string) string {
	i, j, carry, sum := len(a)-1, len(b)-1, 0, ""

	for i >= 0 || j >= 0 {
		temp := processBinaryStrings(BinaryStrings{a, b, &i, &j}, carry)

		carry = temp / divisor
		sum = string(byte(temp%divisor)+'0') + sum
	}

	if carry != 0 {
		sum = "1" + sum
	}

	return sum
}
func processBinaryStrings(strings BinaryStrings, carry int) int {
	temp := carry

	if *strings.i >= 0 {
		if strings.a[*strings.i] == '1' {
			temp++
		}
		*strings.i--
	}

	if *strings.j >= 0 {
		if strings.b[*strings.j] == '1' {
			temp++
		}
		*strings.j--
	}

	return temp
}
