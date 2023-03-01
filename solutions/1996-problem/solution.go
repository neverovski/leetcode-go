package solutions

import (
	"math"
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	SortSlice(properties)

	count := 0
	max := 0

	for _, property := range properties {
		if property[1] < max {
			count += 1
		}

		max = Max(max, property[1])
	}

	return count
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MathMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func SortSlice(properties [][]int) {
	less := func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}

		return properties[i][0] > properties[j][0]
	}

	sort.Slice(properties, less)
}
