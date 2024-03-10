package problem070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, climbStairs(1))
	a.Equal(2, climbStairs(2))
	a.Equal(3, climbStairs(3))
	a.Equal(5, climbStairs(4))
	a.Equal(8, climbStairs(5))
	a.Equal(13, climbStairs(6))
	a.Equal(21, climbStairs(7))
	a.Equal(34, climbStairs(8))
	a.Equal(55, climbStairs(9))
	a.Equal(89, climbStairs(10))
	a.Equal(144, climbStairs(11))
}

func BenchmarkClimbStairs(b *testing.B) {
	n := 30

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		climbStairs(n)
	}
}
