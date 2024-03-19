package problem069

import "math"

const divisor = 2
const initialZ = 100.0
const epsilon = 0.0001

// Using Newton's method (Z + X / Z) / 2
func mySqrt(x int) int {
	z := initialZ

	for {
		root := (z + float64(x)/z) / divisor

		if math.Abs(root-z) < epsilon {
			break
		}

		z = root
	}

	return int(z)
}
