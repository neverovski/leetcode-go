package problem069

import "math"

var INITIAL_Z = 100.0
var EPSILON = 0.0001

// Using Newton's method (Z + X / Z) / 2
func mySqrt(x int) int {
	z := INITIAL_Z

	for {
		root := (z + float64(x)/z) / 2

		if math.Abs(root-z) < EPSILON {
			break
		}

		z = root
	}

	return int(z)
}
