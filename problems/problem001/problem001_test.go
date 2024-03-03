package problem001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	a := assert.New(t)

	a.ElementsMatch([]int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	a.ElementsMatch([]int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	a.ElementsMatch([]int{0, 1}, twoSum([]int{3, 3}, 6))
}
