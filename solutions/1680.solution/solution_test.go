package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  int
	out int
}

var cases = []testCase{
	{1, 1},
	{3, 27},
	{12, 505379714},
	{42, 727837408},
	{9801, 794569523},
}

func BenchmarkConcatenatedBinaryLog2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenatedBinaryLog2(cases[0].in)
	}
}

func BenchmarkConcatenatedBinaryFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenatedBinaryFormatInt(cases[0].in)
	}
}

func BenchmarkConcatenatedBinaryFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenatedBinaryFast(cases[0].in)
	}
}

func TestConcatenatedBinaryLog2(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d, %d", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, concatenatedBinaryLog2(tt.in))
		})
	}
}

func TestConcatenatedBinaryFormatInt(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d, %d", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, concatenatedBinaryFormatInt(tt.in))
		})
	}
}

func TestConcatenatedBinaryFast(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d, %d", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, concatenatedBinaryFast(tt.in))
		})
	}
}
