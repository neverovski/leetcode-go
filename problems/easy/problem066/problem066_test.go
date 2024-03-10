package problem066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	a.Equal([]int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	a.Equal([]int{1, 0}, plusOne([]int{9}))
	a.Equal([]int{1, 0, 0}, plusOne([]int{9, 9}))
	a.Equal([]int{1, 0, 0, 0}, plusOne([]int{9, 9, 9}))
	a.Equal([]int{9, 0, 0, 0}, plusOne([]int{8, 9, 9, 9}))
	a.Equal([]int{7, 9, 0, 0}, plusOne([]int{7, 8, 9, 9}))
}

func BenchmarkPlusOne(b *testing.B) {
	digits := []int{1, 2, 3}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		plusOne(digits)
	}
}
