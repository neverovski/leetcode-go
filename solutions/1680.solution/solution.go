package solutions

import (
	"math"
	"strconv"
)

const MOD = 1e9 + 7

func concatenatedBinaryLog2(n int) int {
	num := 0

	for i := 1; i <= n; i++ {
		length := int(math.Floor(math.Log2(float64(i))) + 1)
		num = (num<<length + i) % MOD
	}

	return num
}

func concatenatedBinaryFormatInt(n int) int {
	num := 0

	for i := 1; i <= n; i++ {
		length := len(strconv.FormatInt(int64(i), 2))
		num = (num<<length + i) % MOD
	}

	return num
}

func concatenatedBinaryFast(n int) int {
	num := 0
	length := 0

	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			length++
		}

		num = (num<<length + i) % MOD
	}

	return num
}
